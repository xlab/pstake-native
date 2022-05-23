package keeper

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/persistenceOne/pstake-native/x/cosmos/types"
)

func (k Keeper) setCosmosValidatorParams(ctx sdk.Context, details types.WeightedAddressAmounts) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.KeyCosmosValidatorSet)
	if store.Has(key) {
		sort.Sort(details)
		bz, err := details.Marshal()
		if err != nil {
			panic(err)
		}
		store.Set(key, bz)
	} else {
		newWeightedAddress := types.NewWeightedAddressAmounts(k.GetParams(ctx).ValidatorSetCosmosChain)
		sort.Sort(newWeightedAddress)
		bz, err := newWeightedAddress.Marshal()
		if err != nil {
			panic(err)
		}
		store.Set(key, bz)
	}
}

func (k Keeper) getCosmosValidatorParams(ctx sdk.Context) (weightedAddrAmt types.WeightedAddressAmounts) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.KeyCosmosValidatorSet))
	err := weightedAddrAmt.Unmarshal(bz)
	if err != nil {
		panic(err)
	}
	return weightedAddrAmt
}

func (k Keeper) updateCosmosValidatorStakingParams(ctx sdk.Context, msgs []sdk.Msg) error {
	uatomDenom, err := k.GetParams(ctx).GetBondDenomOf("uatom")
	if err != nil {
		return err
	}
	totalAmountInDelegateMsgs := sdk.NewInt64Coin(uatomDenom, 0)
	msgsMap := make(map[string]stakingTypes.MsgDelegate, len(msgs))
	for _, msg := range msgs {
		delegateMsg := msg.(*stakingTypes.MsgDelegate)
		totalAmountInDelegateMsgs = totalAmountInDelegateMsgs.Add(delegateMsg.Amount)
		msgsMap[delegateMsg.ValidatorAddress] = *delegateMsg
	}

	k.setTotalDelegatedAmountTillDate(ctx, totalAmountInDelegateMsgs)

	weightedAddressAmounts := k.getCosmosValidatorParams(ctx)
	for _, element := range weightedAddressAmounts {
		if val, ok := msgsMap[element.Address]; ok {
			if element.Denom == val.Amount.Denom {
				element.Amount.Add(val.Amount.Amount)
			}
			// TODO refactor this, difference and ideal delegated amount was deleted.
			//element.IdealDelegatedAmount = sdk.NewCoin(element.IdealDelegatedAmount.Denom,
			//	k.getTotalDelegatedAmountTillDate(ctx).Amount.ToDec().Mul(element.Weight).TruncateInt(),
			//)
			//element.Difference = element.IdealDelegatedAmount.Sub(element.CurrentDelegatedAmount)
		}
	}
	k.setCosmosValidatorParams(ctx, weightedAddressAmounts)
	return nil
	//TODO : Update c token ratio
}

type ValAddressAmount struct {
	Validator sdk.ValAddress
	Amount    sdk.Coin
}

// normalizedWeightedAddressAmounts function takes input as the weighted address amounts
// finds the smallest amount or zero from the array and returns a new array with normalized amounts
func normalizedWeightedAddressAmounts(weightedAddrAmt types.WeightedAddressAmounts) types.WeightedAddressAmounts {
	// Find smallest diff less than zero
	smallestVal := sdk.ZeroInt()
	normalizedDistribution := types.WeightedAddressAmounts{}

	for _, w := range weightedAddrAmt {
		if w.Amount.LT(smallestVal) {
			smallestVal = w.Amount
		}
	}
	// Return early incase the smallest value is zero 
	if smallestVal.Equal(sdk.ZeroInt()) {
		return weightedAddrAmt
	}
	// Normalize based on smallest diff
	for _, w := range weightedAddrAmt {
		normCoin := sdk.NewCoin(w.Denom, w.Amount.Sub(smallestVal))
		normalizedDistribution = append(
			normalizedDistribution,
			types.NewWeightedAddressAmount(w.Address, w.Weight, normCoin),
		)
	}
	return normalizedDistribution
}

func getIdealCurrentDelegations(validatorState types.WeightedAddressAmounts, stakingDenom string) types.WeightedAddressAmounts {
	totalDelegations := validatorState.TotalAmount(stakingDenom)
	curDiffDistribution := types.WeightedAddressAmounts{}
	var idealTokens, curTokens sdk.Int
	for _, valState := range validatorState {
		// Note this can lead to some leaks
		idealTokens = valState.Weight.Mul(totalDelegations.Amount.ToDec()).RoundInt()
		curTokens = valState.Amount

		curDiffDistribution = append(curDiffDistribution, types.WeightedAddressAmount{
			Address: valState.Address,
			Weight: valState.Weight,
			Denom: valState.Denom,
			Amount: idealTokens.Sub(curTokens),
		})
	}
	return curDiffDistribution
}

func divideAmountIntoValidatorSet(valDiff types.WeightedAddressAmounts, coin sdk.Coin) ([]ValAddressAmount, error) {
	if coin.IsZero() {
		return nil, nil
	}

	valAmounts := []ValAddressAmount{}
	for _, w := range valDiff {
		// Skip validators with zero weights
		if w.Weight.IsZero() {
			continue
		}
		// Create val address
		valAddr, err := sdk.ValAddressFromBech32(w.Address)
		if err != nil {
			return nil, err
		}
		if w.Amount.GTE(coin.Amount) {
			valAmounts = append(valAmounts, ValAddressAmount{Validator: valAddr, Amount: coin})
			return valAmounts, nil
		}
		valAmounts = append(valAmounts, ValAddressAmount{Validator: valAddr, Amount: w.Coin()})
		coin = coin.SubAmount(w.Amount)
	}

	// If the remaining amount is not possitive, return early
	if !coin.IsPositive() {
		return valAmounts, nil
	}

	// Divide the remaining amount amongst the validators a/c to weight
	// Note: Maybe there is some slippage due to multiplication
	valAddressMap := types.GetWeightedAddressMap(valDiff)
	for i, valAmt := range valAmounts {
		weight := valAddressMap[valAmt.Validator.String()].Weight
		amt := weight.MulInt(coin.Amount).RoundInt()
		valAmounts[i] = ValAddressAmount{
			Validator: valAmt.Validator,
			Amount: sdk.NewCoin(valAmt.Amount.Denom, valAmt.Amount.Amount.Add(amt)),
		}
	}

	return valAmounts, nil
}

// gives a list of all validators having weighted amount for few and 1uatom for rest in order to auto claim all rewards accumulated in current epoch
func (k Keeper) fetchValidatorsToDelegate(ctx sdk.Context, amount sdk.Coin) ([]ValAddressAmount, error) {
	params := k.GetParams(ctx)

	// Return nil list if amount is less than delegation threshold
	if amount.IsLT(params.DelegationThreshold) {
		return nil, nil
	}

	validatorParams := k.getCosmosValidatorParams(ctx)
	
	curDiffDistribution := getIdealCurrentDelegations(validatorParams, params.StakingDenom)
	curDiffDistribution = normalizedWeightedAddressAmounts(curDiffDistribution)
	
	sort.Sort(curDiffDistribution)

	return divideAmountIntoValidatorSet(curDiffDistribution, amount)
}

// gives a list of validators having weighted amount for few validators
func (k Keeper) fetchValidatorsToUndelegate(ctx sdk.Context, amount sdk.Coin) []ValAddressAmount {
	//TODO : Implement opposite of fetchValidatorsToDelegate
	return nil
}

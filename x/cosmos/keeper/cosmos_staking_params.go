package keeper

import (
	"encoding/json"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/persistenceOne/pstake-native/x/cosmos/types"
)

func (k Keeper) setCosmosValidatorParams(ctx sdk.Context, details types.WeightedAddressAmounts) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.KeyCosmosValidatorSet)
	if store.Has(key) {
		bz, err := details.Sort().Marshal()
		if err != nil {
			panic(err)
		}
		store.Set(key, bz)
	} else {
		newWeightedAddress := ConvertTotypes.WeightedAddressAmounts(k.GetParams(ctx).ValidatorSetCosmosChain)
		bz, err := newWeightedAddress.Sort().Marshal()
		if err != nil {
			panic(err)
		}
		store.Set(key, bz)
	}
}

func (k Keeper) getCosmosValidatorParams(ctx sdk.Context) (types.WeightedAddressAmounts types.WeightedAddressAmounts) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.KeyCosmosValidatorSet))
	err := types.WeightedAddressAmounts.Unmarshal(bz)
	if err != nil {
		panic(err)
	}
	return types.WeightedAddressAmounts
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

	types.WeightedAddressAmounts := k.getCosmosValidatorParams(ctx)
	for _, element := range types.WeightedAddressAmounts {
		if val, ok := msgsMap[element.Address]; ok {
			element.CurrentDelegatedAmount.Add(val.Amount)
			// TODO refactor this, difference and ideal delegated amount was deleted.
			//element.IdealDelegatedAmount = sdk.NewCoin(element.IdealDelegatedAmount.Denom,
			//	k.getTotalDelegatedAmountTillDate(ctx).Amount.ToDec().Mul(element.Weight).TruncateInt(),
			//)
			//element.Difference = element.IdealDelegatedAmount.Sub(element.CurrentDelegatedAmount)
		}
	}
	k.setCosmosValidatorParams(ctx, types.WeightedAddressAmounts)
	return nil
	//TODO : Update c token ratio
}

type ValAddressAndAmountForStakingAndUndelegating struct {
	validator sdk.ValAddress
	amount    sdk.Coin
}

func normalizedTokenDistribution(diffDistribution types.WeightedAddressAmounts) types.WeightedAddressAmounts {
	// Find smallest diff less than zero
	smallestVal := sdk.ZeroInt()
	normalizedDistribution := map[string]sdk.Int{}

	for addr, diff := range diffDistribution {
		normalizedDistribution[addr] = diff
		if diff.LT(smallestVal) {
			smallestVal = diff
		}
	}
	// Return early incase the smallest value is zero 
	if smallestVal.Equal(sdk.ZeroInt()) {
		return diffDistribution
	}
	// Normalize based on smallest diff
	for addr, diff := range diffDistribution {
		normalizedDistribution[addr] = diff.Sub(smallestVal)
	}
	return normalizedDistribution
}

// gives a list of all validators having weighted amount for few and 1uatom for rest in order to auto claim all rewards accumulated in current epoch
func (k Keeper) fetchValidatorsToDelegate(ctx sdk.Context, amount sdk.Coin) []ValAddressAndAmountForStakingAndUndelegating {
	params := k.GetParams(ctx)

	// Return nil list if amount is less than delegation threshold
	if amount.IsLT(params.DelegationThreshold) {
		return nil
	} 

	validatorParams := k.getCosmosValidatorParams(ctx)
	totalDelegations := validatorParams.TotalDelegations(params.StakingDenom)

	curDiffDistribution := types.WeightedAddressAmounts{}

	var idealTokens, curTokens sdk.Int
	for _, valParam := range validatorParams {
		idealTokens = sdk.Int(valParam.Weight.Mul(totalDelegations.Amount.ToDec()))
		curTokens = valParam.CurrentDelegatedAmount.Amount

		curDiffDistribution = append(curDiffDistribution, types.WeightedAddressCosmos{
			Address: valParam.Address,
			Weight: valParam.Weight,
			CurrentDelegatedAmount: sdk.NewCoin(valParam.CurrentDelegatedAmount.Denom, idealTokens.Sub(curTokens)),
		})
		curDiffDistribution[valParam.Address] = idealTokens.Sub(curTokens)
	}
	curDiffDistribution = normalizedTokenDistribution(curDiffDistribution)
	
	// Get the top validators
	sort.Sort(curDiffDistribution)
	

	return nil
}

// gives a list of validators having weighted amount for few validators
func (k Keeper) fetchValidatorsToUndelegate(ctx sdk.Context, amount sdk.Coin) []ValAddressAndAmountForStakingAndUndelegating {
	//TODO : Implement opposite of fetchValidatorsToDelegate
	return nil
}

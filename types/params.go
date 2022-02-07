package types

import (
	"fmt"
	"strings"

	farmingtypes "github.com/cosmosquad-labs/squad/x/farming/types"
	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyBondedBondDenom        = []byte("BondedBondDenom")
	KeyWhitelistedValidators  = []byte("WhitelistedValidators")
	KeyUnstakeFeeRate         = []byte("UnstakeFeeRate")
	KeyCommissionRate         = []byte("CommissionRate")
	KeyMinLiquidStakingAmount = []byte("MinLiquidStakingAmount")

	DefaultBondedBondDenom = "bstake"

	// DefaultUnstakeFeeRate is the default Unstake Fee Rate.
	DefaultUnstakeFeeRate = sdk.NewDecWithPrec(1, 3) // "0.001000000000000000"

	// MinLiquidStakingAmount is the default minimum liquid staking amount.
	DefaultMinLiquidStakingAmount = sdk.NewInt(1000000)

	// Const variables
	RebalancingTrigger = sdk.NewDecWithPrec(1, 3) // "0.001000000000000000"
	RewardTrigger      = sdk.NewDecWithPrec(1, 3) // "0.001000000000000000"

	//LiquidStakingProxyAcc = farmingtypes.DeriveAddress(farmingtypes.AddressType20Bytes, ModuleName, "LiquidStakingProxyAcc")
	LiquidStakingProxyAcc = farmingtypes.DeriveAddress(farmingtypes.AddressType32Bytes, ModuleName, "LiquidStakingProxyAcc")
)

var _ paramstypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns the default liquidstaking module parameters.
func DefaultParams() Params {
	return Params{
		// TODO: btoken denom immutable
		BondedBondDenom:        DefaultBondedBondDenom,
		UnstakeFeeRate:         DefaultUnstakeFeeRate,
		MinLiquidStakingAmount: DefaultMinLiquidStakingAmount,
	}
}

// ParamSetPairs implements paramstypes.ParamSet.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyBondedBondDenom, &p.BondedBondDenom, ValidateBondedBondDenom),
		paramstypes.NewParamSetPair(KeyWhitelistedValidators, &p.WhitelistedValidators, ValidateWhitelistedValidators),
		paramstypes.NewParamSetPair(KeyUnstakeFeeRate, &p.UnstakeFeeRate, validateUnstakeFeeRate),
		paramstypes.NewParamSetPair(KeyMinLiquidStakingAmount, &p.MinLiquidStakingAmount, validateMinLiquidStakingAmount),
	}
}

// String returns a human-readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (p Params) WhitelistedValMap() WhitelistedValMap {
	return GetWhitelistedValMap(p.WhitelistedValidators)
}

// Validate validates parameters.
func (p Params) Validate() error {
	for _, v := range []struct {
		value     interface{}
		validator func(interface{}) error
	}{
		{p.BondedBondDenom, ValidateBondedBondDenom},
		{p.WhitelistedValidators, ValidateWhitelistedValidators},
		{p.UnstakeFeeRate, validateUnstakeFeeRate},
		{p.MinLiquidStakingAmount, validateMinLiquidStakingAmount},
	} {
		if err := v.validator(v.value); err != nil {
			return err
		}
	}
	return nil
}

func ValidateBondedBondDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return fmt.Errorf("bond denom cannot be blank")
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}
	return nil
}

// ValidateWhitelistedValidators validates liquidstaking validator and total weight.
func ValidateWhitelistedValidators(i interface{}) error {
	wvs, ok := i.([]WhitelistedValidator)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	valMap := make(map[string]bool)
	for _, wv := range wvs {
		_, valErr := sdk.ValAddressFromBech32(wv.ValidatorAddress)
		if valErr != nil {
			return valErr
		}

		if wv.TargetWeight.IsNil() {
			return fmt.Errorf("liquidstaking validator target weight must not be nil")
		}

		if !wv.TargetWeight.IsPositive() {
			return fmt.Errorf("liquidstaking validator target weight must be positive: %s", wv.TargetWeight)
		}

		if _, ok := valMap[wv.ValidatorAddress]; ok {
			return fmt.Errorf("liquidstaking validator cannot be duplicated: %s", wv.ValidatorAddress)
		}
		valMap[wv.ValidatorAddress] = true
	}
	return nil
}

func validateUnstakeFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("unstake fee rate must not be nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("unstake fee rate must not be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("unstake fee rate too large: %s", v)
	}

	return nil
}

func validateMinLiquidStakingAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("min liquid staking amount must not be nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("min liquid staking amount must not be negative: %s", v)
	}

	return nil
}
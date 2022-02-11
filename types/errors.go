package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// Sentinel errors for the liquidstaking module.
var (
	ErrActiveLiquidValidatorsNotExists = sdkerrors.Register(ModuleName, 2, "active liquid validators not exists")
	ErrInvalidDenom                    = sdkerrors.Register(ModuleName, 3, "invalid denom")
	ErrInvalidBondDenom                = sdkerrors.Register(ModuleName, 4, "invalid bond denom")
	ErrInvalidBondedBondDenom          = sdkerrors.Register(ModuleName, 5, "invalid liquid bonded bond denom")
	ErrNotImplementedYet               = sdkerrors.Register(ModuleName, 6, "not implemented yet")
	ErrLessThanMinLiquidStakingAmount  = sdkerrors.Register(ModuleName, 7, "staking amount should be over params.min_liquid_staking_amount")
	ErrInvalidBTokenSupply             = sdkerrors.Register(ModuleName, 8, "invalid bonded bon denom supply")
	ErrInvalidActiveLiquidValidators   = sdkerrors.Register(ModuleName, 9, "invalid active liquid validators")
)
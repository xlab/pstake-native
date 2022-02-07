package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/liquidstaking/types"
)

// InitGenesis initializes the liquidstaking module's state from a given genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := types.ValidateGenesis(genState); err != nil {
		panic(err)
	}
	k.SetParams(ctx, genState.Params)

	for _, lv := range genState.LiquidValidators {
		k.SetLiquidValidator(ctx, lv)
	}

	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
}

// ExportGenesis returns the liquidstaking module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)
	liquidValidators := k.GetAllLiquidValidators(ctx)

	return types.NewGenesisState(params, liquidValidators)
}
package keeper_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	squadtypes "github.com/cosmosquad-labs/squad/types"
	farmingtypes "github.com/cosmosquad-labs/squad/x/farming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/x/liquidity/types"
	"github.com/cosmosquad-labs/squad/x/liquidstaking"
	"github.com/cosmosquad-labs/squad/x/liquidstaking/types"
	"github.com/cosmosquad-labs/squad/x/mint"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	simapp "github.com/cosmosquad-labs/squad/app"
	"github.com/cosmosquad-labs/squad/x/liquidstaking/keeper"
)

var (
	BlockTime = 10 * time.Second
)

type KeeperTestSuite struct {
	suite.Suite

	app        *simapp.SquadApp
	ctx        sdk.Context
	keeper     keeper.Keeper
	querier    keeper.Querier
	govHandler govtypes.Handler
	addrs      []sdk.AccAddress
	delAddrs   []sdk.AccAddress
	valAddrs   []sdk.ValAddress
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = simapp.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})
	suite.govHandler = params.NewParamChangeProposalHandler(suite.app.ParamsKeeper)
	stakingParams := stakingtypes.DefaultParams()
	stakingParams.MaxEntries = 200
	stakingParams.MaxValidators = 30
	suite.app.StakingKeeper.SetParams(suite.ctx, stakingParams)

	suite.keeper = suite.app.LiquidStakingKeeper
	suite.querier = keeper.Querier{Keeper: suite.keeper}
	suite.addrs = simapp.AddTestAddrs(suite.app, suite.ctx, 10, sdk.NewInt(1_000_000_000))
	suite.delAddrs = simapp.AddTestAddrs(suite.app, suite.ctx, 10, sdk.NewInt(1_000_000_000))
	suite.valAddrs = simapp.ConvertAddrsToValAddrs(suite.delAddrs)

	suite.ctx = suite.ctx.WithBlockHeight(100).WithBlockTime(squadtypes.MustParseRFC3339("2022-03-01T00:00:00Z"))
	params := suite.keeper.GetParams(suite.ctx)
	params.UnstakeFeeRate = sdk.ZeroDec()
	suite.keeper.SetParams(suite.ctx, params)
	suite.keeper.EndBlocker(suite.ctx)
	// call mint.BeginBlocker for init k.SetLastBlockTime(ctx, ctx.BlockTime())
	mint.BeginBlocker(suite.ctx, suite.app.MintKeeper)
}

func (suite *KeeperTestSuite) CreateValidators(powers []int64) ([]sdk.AccAddress, []sdk.ValAddress) {
	suite.app.BeginBlocker(suite.ctx, abci.RequestBeginBlock{})
	num := len(powers)
	addrs := simapp.AddTestAddrsIncremental(suite.app, suite.ctx, num, sdk.NewInt(1000000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addrs)
	pks := simapp.CreateTestPubKeys(num)

	for i, power := range powers {
		val, err := stakingtypes.NewValidator(valAddrs[i], pks[i], stakingtypes.Description{})
		suite.Require().NoError(err)
		suite.app.StakingKeeper.SetValidator(suite.ctx, val)
		err = suite.app.StakingKeeper.SetValidatorByConsAddr(suite.ctx, val)
		suite.Require().NoError(err)
		suite.app.StakingKeeper.SetNewValidatorByPowerIndex(suite.ctx, val)
		suite.app.StakingKeeper.AfterValidatorCreated(suite.ctx, val.GetOperator())
		newShares, err := suite.app.StakingKeeper.Delegate(suite.ctx, addrs[i], sdk.NewInt(power), stakingtypes.Unbonded, val, true)
		suite.Require().NoError(err)
		suite.Require().Equal(newShares.TruncateInt(), sdk.NewInt(power))
	}

	suite.app.EndBlocker(suite.ctx, abci.RequestEndBlock{})
	return addrs, valAddrs
}

func (suite *KeeperTestSuite) liquidStaking(liquidStaker sdk.AccAddress, stakingAmt sdk.Int) {
	params := suite.keeper.GetParams(suite.ctx)
	btokenBalanceBefore := suite.app.BankKeeper.GetBalance(suite.ctx, liquidStaker, params.BondedBondDenom).Amount
	newShares, bTokenMintAmt, err := suite.keeper.LiquidStaking(suite.ctx, types.LiquidStakingProxyAcc, liquidStaker, sdk.NewCoin(sdk.DefaultBondDenom, stakingAmt))
	btokenBalanceAfter := suite.app.BankKeeper.GetBalance(suite.ctx, liquidStaker, params.BondedBondDenom).Amount
	suite.Require().NoError(err)
	suite.NotEqualValues(newShares, sdk.ZeroDec())
	suite.Require().EqualValues(bTokenMintAmt, btokenBalanceAfter.Sub(btokenBalanceBefore))
}

func (s *KeeperTestSuite) advanceHeight(height int, withEndBlock bool) {
	feeCollector := s.app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	for i := 0; i < height; i++ {
		s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(s.ctx.BlockTime().Add(BlockTime))
		mint.BeginBlocker(s.ctx, s.app.MintKeeper)
		feeCollectorBalance := s.app.BankKeeper.GetAllBalances(s.ctx, feeCollector)
		rewardsToBeDistributed := feeCollectorBalance.AmountOf(sdk.DefaultBondDenom)

		// mimic AllocateTokens, get rewards from feeCollector, AllocateTokensToValidator, add remaining to feePool
		err := s.app.BankKeeper.SendCoinsFromModuleToModule(s.ctx, authtypes.FeeCollectorName, distrtypes.ModuleName, feeCollectorBalance)
		s.Require().NoError(err)
		totalRewards := sdk.ZeroDec()
		totalPower := int64(0)
		s.app.StakingKeeper.IterateBondedValidatorsByPower(s.ctx, func(index int64, validator stakingtypes.ValidatorI) (stop bool) {
			consPower := validator.GetConsensusPower(s.app.StakingKeeper.PowerReduction(s.ctx))
			totalPower = totalPower + consPower
			return false
		})
		s.app.StakingKeeper.IterateBondedValidatorsByPower(s.ctx, func(index int64, validator stakingtypes.ValidatorI) (stop bool) {
			consPower := validator.GetConsensusPower(s.app.StakingKeeper.PowerReduction(s.ctx))
			powerFraction := sdk.NewDec(consPower).QuoTruncate(sdk.NewDec(totalPower))
			reward := rewardsToBeDistributed.ToDec().MulTruncate(powerFraction)
			s.app.DistrKeeper.AllocateTokensToValidator(s.ctx, validator, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: reward}})
			totalRewards = totalRewards.Add(reward)
			return false
		})
		remaining := rewardsToBeDistributed.ToDec().Sub(totalRewards)
		s.Require().False(remaining.GT(sdk.NewDec(1)))
		feePool := s.app.DistrKeeper.GetFeePool(s.ctx)
		feePool.CommunityPool = feePool.CommunityPool.Add(sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: remaining}}...)
		s.app.DistrKeeper.SetFeePool(s.ctx, feePool)
		staking.BeginBlocker(s.ctx, *s.app.StakingKeeper)
		staking.EndBlocker(s.ctx, *s.app.StakingKeeper)
		if withEndBlock {
			liquidstaking.EndBlocker(s.ctx, s.app.LiquidStakingKeeper)
		}
	}
}

func (s *KeeperTestSuite) fundAddr(addr sdk.AccAddress, amt sdk.Coins) {
	err := s.app.BankKeeper.MintCoins(s.ctx, liquiditytypes.ModuleName, amt)
	s.Require().NoError(err)
	err = s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, liquiditytypes.ModuleName, addr, amt)
	s.Require().NoError(err)
}

// liquidity module keeper utils for liquid staking combine test

func (s *KeeperTestSuite) createPair(creator sdk.AccAddress, baseCoinDenom, quoteCoinDenom string, fund bool) liquiditytypes.Pair {
	params := s.app.LiquidityKeeper.GetParams(s.ctx)
	if fund {
		s.fundAddr(creator, params.PairCreationFee)
	}
	pair, err := s.app.LiquidityKeeper.CreatePair(s.ctx, liquiditytypes.NewMsgCreatePair(creator, baseCoinDenom, quoteCoinDenom))
	s.Require().NoError(err)
	return pair
}

func (s *KeeperTestSuite) createPool(creator sdk.AccAddress, pairId uint64, depositCoins sdk.Coins, fund bool) liquiditytypes.Pool {
	params := s.app.LiquidityKeeper.GetParams(s.ctx)
	if fund {
		s.fundAddr(creator, depositCoins.Add(params.PoolCreationFee...))
	}
	pool, err := s.app.LiquidityKeeper.CreatePool(s.ctx, liquiditytypes.NewMsgCreatePool(creator, pairId, depositCoins))
	s.Require().NoError(err)
	return pool
}

// farming module keeper utils for liquid staking combine test

func (s *KeeperTestSuite) AdvanceEpoch() {
	err := s.app.FarmingKeeper.AdvanceEpoch(s.ctx)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) CreateFixedAmountPlan(farmingPoolAcc sdk.AccAddress, stakingCoinWeightsMap map[string]string, epochAmountMap map[string]int64) {
	stakingCoinWeights := sdk.NewDecCoins()
	for denom, weight := range stakingCoinWeightsMap {
		stakingCoinWeights = stakingCoinWeights.Add(sdk.NewDecCoinFromDec(denom, sdk.MustNewDecFromStr(weight)))
	}

	epochAmount := sdk.NewCoins()
	for denom, amount := range epochAmountMap {
		epochAmount = epochAmount.Add(sdk.NewInt64Coin(denom, amount))
	}

	msg := farmingtypes.NewMsgCreateFixedAmountPlan(
		fmt.Sprintf("plan%d", s.app.FarmingKeeper.GetGlobalPlanId(s.ctx)+1),
		farmingPoolAcc,
		stakingCoinWeights,
		farmingtypes.ParseTime("0001-01-01T00:00:00Z"),
		farmingtypes.ParseTime("9999-12-31T00:00:00Z"),
		epochAmount,
	)
	_, err := s.app.FarmingKeeper.CreateFixedAmountPlan(s.ctx, msg, farmingPoolAcc, farmingPoolAcc, farmingtypes.PlanTypePublic)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) Stake(farmerAcc sdk.AccAddress, amt sdk.Coins) {
	err := s.app.FarmingKeeper.Stake(s.ctx, farmerAcc, amt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) Unstake(farmerAcc sdk.AccAddress, amt sdk.Coins) {
	err := s.app.FarmingKeeper.Unstake(s.ctx, farmerAcc, amt)
	s.Require().NoError(err)
}

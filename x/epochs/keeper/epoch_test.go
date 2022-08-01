package keeper_test

import (
	"time"

	"github.com/persistenceOne/pstake-native/x/epochs/types"
)

func (suite *KeeperTestSuite) TestEpochLifeCycle() {
	suite.SetupTest()

	epochInfo := types.EpochInfo{
		Identifier:            "monthly",
		StartTime:             time.Time{},
		Duration:              time.Hour * 24 * 30,
		CurrentEpoch:          0,
		CurrentEpochStartTime: time.Time{},
		EpochCountingStarted:  false,
	}
	suite.app.EpochsKeeper.SetEpochInfo(suite.ctx, epochInfo)
	epochInfoSaved := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, "monthly")
	suite.Require().Equal(epochInfo, epochInfoSaved)

	allEpochs := suite.app.EpochsKeeper.AllEpochInfos(suite.ctx)
	suite.Require().Len(allEpochs, 4)
	suite.Require().Equal(allEpochs[0].Identifier, "monthly") // alphabetical order
	suite.Require().Equal(allEpochs[1].Identifier, "reward")
	suite.Require().Equal(allEpochs[2].Identifier, "stake")
	suite.Require().Equal(allEpochs[3].Identifier, "undelegate")
}

package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"

	epochtypes "github.com/Stride-Labs/stride/v22/x/epochs/types"
	recordtypes "github.com/Stride-Labs/stride/v22/x/records/types"
	stakeibctypes "github.com/Stride-Labs/stride/v22/x/stakeibc/types"
	"github.com/Stride-Labs/stride/v22/x/staketia/keeper"
	oldtypes "github.com/Stride-Labs/stride/v22/x/staketia/legacytypes"
	"github.com/Stride-Labs/stride/v22/x/staketia/types"
)

func (s *KeeperTestSuite) TestUpdateStakeibcHostZone() {
	// Create deposit records with amounts 100 and 200 respectively
	delegationRecords := []types.DelegationRecord{
		{Id: 1, Status: types.TRANSFER_IN_PROGRESS, NativeAmount: sdk.NewInt(100)},
		{Id: 2, Status: types.DELEGATION_QUEUE, NativeAmount: sdk.NewInt(200)},
	}
	for _, delegationRecord := range delegationRecords {
		s.App.StaketiaKeeper.SetDelegationRecord(s.Ctx, delegationRecord)
	}

	// Create a host zone with a delegated balance of 1000
	redemptionRate := sdk.NewDec(2)
	legacyHostZone := oldtypes.HostZone{
		RedemptionRate:   redemptionRate,
		DelegatedBalance: sdk.NewInt(1_000),
	}
	stakeibcHostZone := stakeibctypes.HostZone{
		ChainId: types.CelestiaChainId,
	}
	s.App.StaketiaKeeper.SetLegacyHostZone(s.Ctx, legacyHostZone)
	s.App.StakeibcKeeper.SetHostZone(s.Ctx, stakeibcHostZone)

	// The expected stakeibc host zone should have total delegations
	// equal to 1000 + 100 + 200 = 1300
	expectedStakeibcTotalDelegations := sdkmath.NewInt(1_000 + 100 + 200)

	// Call the update host zone function and confirm against expectations
	actualStakeibcHostZone, err := s.App.StaketiaKeeper.UpdateStakeibcHostZone(s.Ctx, legacyHostZone)
	s.Require().NoError(err, "no error expected when updating host zone")

	s.Require().Equal(types.CelestiaChainId, actualStakeibcHostZone.ChainId, "chain ID")
	s.Require().Equal(expectedStakeibcTotalDelegations, actualStakeibcHostZone.TotalDelegations, "total delegations")
	s.Require().Equal(redemptionRate, actualStakeibcHostZone.RedemptionRate, "redemption rate")

	// Remove the host zone and try again, it should fail
	s.App.StakeibcKeeper.RemoveHostZone(s.Ctx, types.CelestiaChainId)
	_, err = s.App.StaketiaKeeper.UpdateStakeibcHostZone(s.Ctx, legacyHostZone)
	s.Require().ErrorContains(err, "celestia host zone not found")
}

func (s *KeeperTestSuite) TestMigrateProtocolOwnedAccounts() {
	// Create deposit accounts across both modules
	staketiaDepositAccount := s.TestAccs[0]
	stakeibcDepositAccount := s.TestAccs[1]

	// Get the respective fee module accounts for both modules
	staketiaFeeModuleName := types.FeeAddress
	stakeibcFeeAddress := s.App.AccountKeeper.GetModuleAddress(stakeibctypes.RewardCollectorName)

	// Set the addresses on the respective host zones
	staketiaHostZone := oldtypes.HostZone{
		DepositAddress:      staketiaDepositAccount.String(),
		NativeTokenIbcDenom: HostIBCDenom,
	}
	stakeibcHostZone := stakeibctypes.HostZone{
		ChainId:        types.CelestiaChainId,
		DepositAddress: stakeibcDepositAccount.String(),
	}

	// Create a deposit record that will be modified
	s.App.RecordsKeeper.SetDepositRecord(s.Ctx, recordtypes.DepositRecord{
		Id:         1,
		Amount:     sdkmath.ZeroInt(),
		HostZoneId: types.CelestiaChainId,
	})

	// Fund the deposit and fee account on staketia
	denom := staketiaHostZone.NativeTokenIbcDenom
	expectedDepositBalance := sdkmath.NewInt(1000)
	expectedFeeBalance := sdkmath.NewInt(2000)

	s.FundAccount(staketiaDepositAccount, sdk.NewCoin(denom, expectedDepositBalance))
	s.FundModuleAccount(staketiaFeeModuleName, sdk.NewCoin(denom, expectedFeeBalance))

	// Call the migration function to transfer to stakeibc
	err := s.App.StaketiaKeeper.MigrateProtocolOwnedAccounts(s.Ctx, staketiaHostZone, stakeibcHostZone)
	s.Require().NoError(err, "no error expected when migrating accounts")

	// Check that the stakeibc accounts are now funded
	actualDepositBalance := s.App.BankKeeper.GetBalance(s.Ctx, stakeibcDepositAccount, denom)
	s.Require().Equal(expectedDepositBalance.Int64(), actualDepositBalance.Amount.Int64(), "deposit balance")

	actualFeeBalance := s.App.BankKeeper.GetBalance(s.Ctx, stakeibcFeeAddress, denom)
	s.Require().Equal(expectedFeeBalance.Int64(), actualFeeBalance.Amount.Int64(), "fee balance")

	// Confirm that the deposit record was incremented
	depositRecords := s.App.RecordsKeeper.GetAllDepositRecord(s.Ctx)
	s.Require().Len(depositRecords, 1, "deposit record should have been created")
	s.Require().Equal(expectedDepositBalance.Int64(), depositRecords[0].Amount.Int64(), "deposit record")

	// Create a second deposit record and try to call the migration again, it should fail
	s.App.RecordsKeeper.SetDepositRecord(s.Ctx, recordtypes.DepositRecord{
		Id:         2,
		HostZoneId: types.CelestiaChainId,
	})
	err = s.App.StaketiaKeeper.MigrateProtocolOwnedAccounts(s.Ctx, staketiaHostZone, stakeibcHostZone)
	s.Require().ErrorContains(err, "there should only be one celestia deposit record")
}

func (s *KeeperTestSuite) TestInitiateMigration() {
	// Create a transfer channel (which will create a connection)
	s.CreateTransferChannel(types.CelestiaChainId)

	staketiaDepositAccount := s.TestAccs[0]
	staketiaFeeModuleName := types.FeeAddress

	// Fund the staketia deposit and fee accounts
	depositBalance := sdkmath.NewInt(1000)
	feeBalance := sdkmath.NewInt(2000)
	s.FundAccount(staketiaDepositAccount, sdk.NewCoin(HostIBCDenom, depositBalance))
	s.FundModuleAccount(staketiaFeeModuleName, sdk.NewCoin(HostIBCDenom, feeBalance))

	// Store the legacy host zone
	legacyHostZone := oldtypes.HostZone{
		ChainId:             types.CelestiaChainId,
		DepositAddress:      staketiaDepositAccount.String(),
		NativeTokenDenom:    HostNativeDenom,
		NativeTokenIbcDenom: HostIBCDenom,
		TransferChannelId:   ibctesting.FirstChannelID,
		MinRedemptionRate:   sdk.MustNewDecFromStr("0.90"),
		MaxRedemptionRate:   sdk.MustNewDecFromStr("1.5"),
		RedemptionRate:      sdk.MustNewDecFromStr("1.2"),
		DelegatedBalance:    sdk.NewInt(1000),
	}
	s.App.StaketiaKeeper.SetLegacyHostZone(s.Ctx, legacyHostZone)

	// Create a delegation record that will be used in the delegated balance migration
	delegationRecord := types.DelegationRecord{
		Id:           1,
		Status:       types.DELEGATION_QUEUE,
		NativeAmount: sdk.NewInt(100),
	}
	s.App.StaketiaKeeper.SetDelegationRecord(s.Ctx, delegationRecord)
	expectedTotalDelegations := legacyHostZone.DelegatedBalance.Add(delegationRecord.NativeAmount)

	// Create epoch trackers and EURs which are needed for the stakeibc registration
	s.App.StakeibcKeeper.SetEpochTracker(s.Ctx, stakeibctypes.EpochTracker{
		EpochIdentifier: epochtypes.DAY_EPOCH,
		EpochNumber:     uint64(1),
	})
	s.App.StakeibcKeeper.SetEpochTracker(s.Ctx, stakeibctypes.EpochTracker{
		EpochIdentifier: epochtypes.STRIDE_EPOCH,
		EpochNumber:     uint64(1),
	})
	epochUnbondingRecord := recordtypes.EpochUnbondingRecord{
		EpochNumber:        uint64(1),
		HostZoneUnbondings: []*recordtypes.HostZoneUnbonding{},
	}
	s.App.RecordsKeeper.SetEpochUnbondingRecord(s.Ctx, epochUnbondingRecord)

	// Call the migration function to register with stakeibc
	// Before we call it, temporarily update the variable to be connection-0 to match the above
	// and then set it back after the function call for other tests that use it
	mainnetConnectionId := types.CelestiaConnectionId
	types.CelestiaConnectionId = ibctesting.FirstConnectionID

	err := keeper.InitiateMigration(s.App.StaketiaKeeper, s.Ctx)
	types.CelestiaConnectionId = mainnetConnectionId
	s.Require().NoError(err, "no error expected during migration")

	// Confirm the new host zone
	hostZone, found := s.App.StakeibcKeeper.GetHostZone(s.Ctx, types.CelestiaChainId)
	s.Require().True(found, "stakeibc host zone should have been created")

	s.Require().Equal(legacyHostZone.TransferChannelId, hostZone.TransferChannelId, "transfer channel ID")
	s.Require().Equal(legacyHostZone.NativeTokenDenom, hostZone.HostDenom, "native denom")
	s.Require().Equal(legacyHostZone.NativeTokenIbcDenom, hostZone.IbcDenom, "ibc denom")

	s.Require().Equal(legacyHostZone.RedemptionRate, hostZone.RedemptionRate, "redemption rate")
	s.Require().Equal(legacyHostZone.MinRedemptionRate, hostZone.MinRedemptionRate, "min redemption rate")
	s.Require().Equal(legacyHostZone.MaxRedemptionRate, hostZone.MaxRedemptionRate, "max redemption rate")

	s.Require().Equal(ibctesting.FirstConnectionID, hostZone.ConnectionId, "connection ID")
	s.Require().Equal(types.CelestiaBechPrefix, hostZone.Bech32Prefix, "bech prefix")
	s.Require().Equal(uint64(types.CelestiaUnbondingPeriodDays), hostZone.UnbondingPeriod, "unbonding period")

	s.Require().False(hostZone.RedemptionsEnabled, "redemptions enabled")
	s.Require().Equal(expectedTotalDelegations, hostZone.TotalDelegations, "total delegations")

	// Confirm balances were transferred
	stakeibcDepositAccount := sdk.MustAccAddressFromBech32(hostZone.DepositAddress)
	actualDepositBalance := s.App.BankKeeper.GetBalance(s.Ctx, stakeibcDepositAccount, HostIBCDenom)
	s.Require().Equal(depositBalance.Int64(), actualDepositBalance.Amount.Int64(), "deposit balance transfer")

	stakeibcFeeAddress := s.App.AccountKeeper.GetModuleAddress(stakeibctypes.RewardCollectorName)
	actualFeeBalance := s.App.BankKeeper.GetBalance(s.Ctx, stakeibcFeeAddress, HostIBCDenom)
	s.Require().Equal(feeBalance.Int64(), actualFeeBalance.Amount.Int64(), "fee balance transfer")

	// Confirm a deposit record was created with the deposit amount
	depositRecords := s.App.RecordsKeeper.GetAllDepositRecord(s.Ctx)
	s.Require().Len(depositRecords, 1, "deposit record should have been created")
	s.Require().Equal(depositBalance.Int64(), depositRecords[0].Amount.Int64(), "deposit record")
}
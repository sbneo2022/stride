package v24

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	recordskeeper "github.com/Stride-Labs/stride/v23/x/records/keeper"
	recordstypes "github.com/Stride-Labs/stride/v23/x/records/types"
	stakeibckeeper "github.com/Stride-Labs/stride/v23/x/stakeibc/keeper"
	staketiakeeper "github.com/Stride-Labs/stride/v23/x/staketia/keeper"
)

var (
	UpgradeName = "v24"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v23
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	recordsKeeper recordskeeper.Keeper,
	stakeibcKeeper stakeibckeeper.Keeper,
	staketiaKeeper staketiakeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("Starting upgrade v24...")

		// Migrate data structures
		MigrateHostZones(ctx, stakeibcKeeper)
		MigrateEpochUnbondingRecords(ctx, recordsKeeper)

		// Migrate staketia to stakeibc
		if err := staketiakeeper.InitiateMigration(staketiaKeeper, ctx); err != nil {
			return vm, errorsmod.Wrapf(err, "unable to migrate staketia to stakeibc")
		}

		// TODO: add celestia validator set

		ctx.Logger().Info("Running module migrations...")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}

// Migrate host zones to accomodate the staketia migration changes, adding a
// redemptions enabled field to each host zone
func MigrateHostZones(ctx sdk.Context, k stakeibckeeper.Keeper) {
	ctx.Logger().Info("Migrating host zones...")

	for _, hostZone := range k.GetAllHostZone(ctx) {
		hostZone.RedemptionsEnabled = true
		k.SetHostZone(ctx, hostZone)
	}
}

// Migrates the deposit records to set the DelegationTxsInProgress field
// which should be 1 if the status was DELEGATION_IN_PROGRESS, and 0 otherwise
func MigrateDepositRecords(ctx sdk.Context, k recordskeeper.Keeper) {
	ctx.Logger().Info("Migrating deposit records...")

	for _, depositRecord := range k.GetAllDepositRecord(ctx) {
		if depositRecord.Status == recordstypes.DepositRecord_DELEGATION_IN_PROGRESS {
			depositRecord.DelegationTxsInProgress = 1
		} else {
			depositRecord.DelegationTxsInProgress = 0
		}
		k.SetDepositRecord(ctx, depositRecord)
	}
}

// Migrates a single host zone unbonding record to add the new fields: StTokensToBurn,
// NativeTokensToUnbond, and ClaimableNativeTokens
//
// If the record is in status: UNBONDING_QUEUE, EXIT_TRANSFER_QUEUE, or EXIT_TRANSFER_IN_PROGRESS,
// set stTokensToBurn, NativeTokensToUnbond, and ClaimableNativeTokens all to 0
//
// If the record is in status: UNBONDING_IN_PROGRESS
// set StTokensToBurn to the value of StTokenAmount, NativeTokensToUnbond to the value of NativeTokenAmount,
// and ClaimableNativeTokens to 0
//
// If the record is in status CLAIMABLE,
// set StTokensToBurn and NativeTokensToUnbond to 0, and set ClaimableNativeTokens to the value of NativeTokenAmount
//
// If the record is in status UNBONDING_IN_PROGRESS, we need to also set UndelegationTxsInProgress to 1;
// otherwise, it should be set to 0
func MigrateHostZoneUnbondingRecords(hostZoneUnbonding *recordstypes.HostZoneUnbonding) *recordstypes.HostZoneUnbonding {
	if hostZoneUnbonding.Status == recordstypes.HostZoneUnbonding_UNBONDING_QUEUE ||
		hostZoneUnbonding.Status == recordstypes.HostZoneUnbonding_EXIT_TRANSFER_QUEUE ||
		hostZoneUnbonding.Status == recordstypes.HostZoneUnbonding_EXIT_TRANSFER_IN_PROGRESS {

		hostZoneUnbonding.StTokensToBurn = sdkmath.ZeroInt()
		hostZoneUnbonding.NativeTokensToUnbond = sdkmath.ZeroInt()
		hostZoneUnbonding.ClaimableNativeTokens = sdkmath.ZeroInt()
		hostZoneUnbonding.UndelegationTxsInProgress = 0

	} else if hostZoneUnbonding.Status == recordstypes.HostZoneUnbonding_UNBONDING_IN_PROGRESS {
		hostZoneUnbonding.StTokensToBurn = hostZoneUnbonding.StTokenAmount
		hostZoneUnbonding.NativeTokensToUnbond = hostZoneUnbonding.NativeTokenAmount
		hostZoneUnbonding.ClaimableNativeTokens = sdkmath.ZeroInt()
		hostZoneUnbonding.UndelegationTxsInProgress = 1

	} else if hostZoneUnbonding.Status == recordstypes.HostZoneUnbonding_CLAIMABLE {
		hostZoneUnbonding.StTokensToBurn = sdkmath.ZeroInt()
		hostZoneUnbonding.NativeTokensToUnbond = sdkmath.ZeroInt()
		hostZoneUnbonding.ClaimableNativeTokens = hostZoneUnbonding.NativeTokenAmount
		hostZoneUnbonding.UndelegationTxsInProgress = 0
	}

	return hostZoneUnbonding
}

// Migrate epoch unbonding records to accomodate the batched undelegations code changes,
// adding the new accounting fields to the host zone unbonding records
func MigrateEpochUnbondingRecords(ctx sdk.Context, k recordskeeper.Keeper) {
	ctx.Logger().Info("Migrating epoch unbonding records...")

	for _, epochUnbondingRecord := range k.GetAllEpochUnbondingRecord(ctx) {
		for i, oldHostZoneUnbondingRecord := range epochUnbondingRecord.HostZoneUnbondings {
			updatedHostZoneUnbondingRecord := MigrateHostZoneUnbondingRecords(oldHostZoneUnbondingRecord)
			epochUnbondingRecord.HostZoneUnbondings[i] = updatedHostZoneUnbondingRecord
		}
		k.SetEpochUnbondingRecord(ctx, epochUnbondingRecord)
	}
}
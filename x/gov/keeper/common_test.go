package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	TestProposal = types.NewTextProposal("Test", "description")
)

func createValidators(t *testing.T, ctx sdk.Context, app *simapp.SimApp, powers []int64) ([]sdk.AccAddress, []sdk.ValAddress) {
	addrs := simapp.AddTestAddrsIncremental(app, ctx, 5, sdk.NewInt(30000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addrs)
	pks := simapp.CreateTestPubKeys(5)
	cdc := simapp.MakeTestEncodingConfig().Marshaler

	app.StakingKeeper = stakingkeeper.NewKeeper(
		cdc,
		app.GetKey(stakingtypes.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)

	randomEthAddress1, err := teststaking.RandomEthAddress()
	require.NoError(t, err)
	val1, err := stakingtypes.NewValidator(valAddrs[0], pks[0], stakingtypes.Description{}, sdk.AccAddress(pks[0].Address()), *randomEthAddress1)
	require.NoError(t, err)

	randomEthAddress2, err := teststaking.RandomEthAddress()
	require.NoError(t, err)
	val2, err := stakingtypes.NewValidator(valAddrs[1], pks[1], stakingtypes.Description{}, sdk.AccAddress(pks[1].Address()), *randomEthAddress2)
	require.NoError(t, err)

	randomEthAddress3, err := teststaking.RandomEthAddress()
	require.NoError(t, err)
	val3, err := stakingtypes.NewValidator(valAddrs[2], pks[2], stakingtypes.Description{}, sdk.AccAddress(pks[2].Address()), *randomEthAddress3)
	require.NoError(t, err)

	app.StakingKeeper.SetValidator(ctx, val1)
	app.StakingKeeper.SetValidator(ctx, val2)
	app.StakingKeeper.SetValidator(ctx, val3)
	app.StakingKeeper.SetValidatorByConsAddr(ctx, val1)
	app.StakingKeeper.SetValidatorByConsAddr(ctx, val2)
	app.StakingKeeper.SetValidatorByConsAddr(ctx, val3)
	app.StakingKeeper.SetNewValidatorByPowerIndex(ctx, val1)
	app.StakingKeeper.SetNewValidatorByPowerIndex(ctx, val2)
	app.StakingKeeper.SetNewValidatorByPowerIndex(ctx, val3)

	_, _ = app.StakingKeeper.Delegate(ctx, addrs[0], app.StakingKeeper.TokensFromConsensusPower(ctx, powers[0]), stakingtypes.Unbonded, val1, true)
	_, _ = app.StakingKeeper.Delegate(ctx, addrs[1], app.StakingKeeper.TokensFromConsensusPower(ctx, powers[1]), stakingtypes.Unbonded, val2, true)
	_, _ = app.StakingKeeper.Delegate(ctx, addrs[2], app.StakingKeeper.TokensFromConsensusPower(ctx, powers[2]), stakingtypes.Unbonded, val3, true)

	_ = staking.EndBlocker(ctx, app.StakingKeeper)

	return addrs, valAddrs
}

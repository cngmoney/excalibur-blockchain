package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "excalibur/testutil/keeper"
	"excalibur/x/excalibur/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ExcaliburKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

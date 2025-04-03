package excalibur_test

import (
	"testing"

	keepertest "excalibur/testutil/keeper"
	"excalibur/testutil/nullify"
	excalibur "excalibur/x/excalibur/module"
	"excalibur/x/excalibur/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExcaliburKeeper(t)
	excalibur.InitGenesis(ctx, k, genesisState)
	got := excalibur.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

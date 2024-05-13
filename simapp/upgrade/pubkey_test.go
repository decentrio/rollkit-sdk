package upgrade_test

import (
	"testing"

	"github.com/decentrio/rollkit-sdk/simapp/upgrade"
	"github.com/stretchr/testify/require"
)

func TestConvertPubKey(t *testing.T) {
	pubKey, err := upgrade.GetSequencerEd25519Pubkey()
	require.NoError(t, err)
	require.Equal(t, pubKey.Address().String(), "E49ADA38175BA34DD0AEF81D531E51C0B110E817")
}

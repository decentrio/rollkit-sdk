package upgrade_test

import (
	"testing"

	"github.com/decentrio/rollkit-sdk/simapp/upgrade"
	"github.com/stretchr/testify/require"
)

func TestConvertPubKey(t *testing.T) {
	const SequencerConsensusPubkeyBase64 = "R+YT9Fz+w5+bCdPUW+IydtOYlTFcS7Irovxu/Xut2S4="

	pubKey, err := upgrade.GetSequencerEd25519Pubkey(SequencerConsensusPubkeyBase64)
	require.NoError(t, err)
	require.Equal(t, pubKey.Address().String(), "DCA26A9DCC0380A05B17352EB3036392F37AFB38")
}

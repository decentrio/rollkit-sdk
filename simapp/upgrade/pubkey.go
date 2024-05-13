package upgrade

import (
	b64 "encoding/base64"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

const SequencerConsensusPubkeyBase64 = "GIBleow/Cud8aXLKj9KNgNqhzieheLTN/HfUETYzoUc="

func GetSequencerEd25519Pubkey() (cryptotypes.PubKey, error) {
	sDec, err := b64.StdEncoding.DecodeString(SequencerConsensusPubkeyBase64)
	if err != nil {
		return nil, err
	}
	return &ed25519.PubKey{Key: sDec}, nil
}

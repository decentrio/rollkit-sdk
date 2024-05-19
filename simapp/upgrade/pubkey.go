package upgrade

import (
	b64 "encoding/base64"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func GetSequencerEd25519Pubkey(base64Pubkey string) (cryptotypes.PubKey, error) {
	sDec, err := b64.StdEncoding.DecodeString(base64Pubkey)
	if err != nil {
		return nil, err
	}
	return &ed25519.PubKey{Key: sDec}, nil
}

package types

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// NewGenesisState creates a new GenesisState instance
func NewGenesisState(params Params, sequencer []Sequencer) *GenesisState {
	return &GenesisState{
		Params:     params,
		Sequencers: sequencer,
	}
}

// DefaultGenesisState gets the raw genesis raw message for testing
func DefaultGenesisState() *GenesisState {
	pubkey := secp256k1.GenPrivKey().PubKey()
	pkAny, err := codectypes.NewAnyWithValue(pubkey)
	if err != nil {
		return &GenesisState{}
	}

	return &GenesisState{
		Sequencers: []Sequencer{
			{
				Name:            "default",
				ConsensusPubkey: pkAny,
			},
		},
		Params: Params{},
	}
}

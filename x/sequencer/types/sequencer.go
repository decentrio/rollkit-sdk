package types

import (
	"cosmossdk.io/errors"
	cmtprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (seq Sequencer) TmConsPublicKey() (cmtprotocrypto.PublicKey, error) {
	pk, ok := seq.ConsensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return cmtprotocrypto.PublicKey{}, errors.Wrapf(sdkerrors.ErrInvalidType, "expecting cryptotypes.PubKey, got %T", pk)
	}

	tmPk, err := cryptocodec.ToCmtProtoPublicKey(pk)
	if err != nil {
		return cmtprotocrypto.PublicKey{}, err
	}
	return tmPk, nil

}

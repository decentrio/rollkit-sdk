package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/decentrio/rollkit-sdk/x/sequencer/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) ChangeSequencers(goCtx context.Context, msg *types.MsgChangeSequencers) (*types.MsgChangeSequencersResponse, error) {
	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if len(msg.Sequencers) != 0 {
		newSequencer := msg.Sequencers[0]
		k.Sequencer.Set(ctx, newSequencer)
		k.NextSequencerChangeHeight.Set(ctx, ctx.BlockHeight())
	}
	return &types.MsgChangeSequencersResponse{}, nil
}

func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Params.Set(ctx, msg.Params); err != nil {
		return nil, err
	}
	
	return &types.MsgUpdateParamsResponse{}, nil
}

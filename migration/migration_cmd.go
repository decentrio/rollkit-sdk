package migration

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	dbm "github.com/cometbft/cometbft-db"
	cometbftcmd "github.com/cometbft/cometbft/cmd/cometbft/commands"
	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/os"
	"github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/store"
	cometbfttypes "github.com/cometbft/cometbft/types"
	rollkitstore "github.com/rollkit/rollkit/store"
	rollkittypes "github.com/rollkit/rollkit/types"
	"github.com/spf13/cobra"
)

// TODO: testing
// MigrateToRollkitCmd returns a command that migrates the data from the comnettBFT chain to rollup
func MigrateToRollkitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollup-migration ",
		Short: "Migrate the data from the comnettBFT chain to rollup",
		Long:  "Migrate the data from the comnettBFT chain to rollup",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := cometbftcmd.ParseConfig(cmd)
			if err != nil {
				return err
			}

			blockStore, stateStore, err := loadStateAndBlockStore(config)
			if err != nil {
				return err
			}

			cometBFTstate, err := stateStore.Load()
			if err != nil {
				return err
			}
			height := cometBFTstate.LastBlockHeight
			block := blockStore.LoadBlock(height)
			rollkitCommit := rollkitCommitFromCometBFTCommit(*block.LastCommit)

			rollkitStore, err := loadRollkitStateStore(config.RootDir, config.DBPath)
			if err != nil {
				return err
			}

			rollkitState, err := rollkitStateFromCometBFTState(cometBFTstate)
			if err != nil {
				return err
			}

			err = rollkitStore.UpdateState(context.Background(), rollkitState)
			if err != nil {
				return err
			}

			rollkitBlock := rollkitBlockFromCometBFTBlock(block, *block.LastCommit, rollkitState.Validators)
			err = rollkitStore.SaveBlock(context.Background(), &rollkitBlock, &rollkitCommit)
			if err != nil {
				return err
			}

			log.Println("Migration completed successfully")
			return nil
		},
	}
	return cmd
}

func loadStateAndBlockStore(config *cfg.Config) (*store.BlockStore, state.Store, error) {
	dbType := dbm.BackendType(config.DBBackend)

	if !os.FileExists(filepath.Join(config.DBDir(), "blockstore.db")) {
		return nil, nil, fmt.Errorf("no blockstore found in %v", config.DBDir())
	}

	// Get BlockStore
	blockStoreDB, err := dbm.NewDB("blockstore", dbType, config.DBDir())
	if err != nil {
		return nil, nil, err
	}
	blockStore := store.NewBlockStore(blockStoreDB)

	if !os.FileExists(filepath.Join(config.DBDir(), "state.db")) {
		return nil, nil, fmt.Errorf("no statestore found in %v", config.DBDir())
	}

	// Get StateStore
	stateDB, err := dbm.NewDB("state", dbType, config.DBDir())
	if err != nil {
		return nil, nil, err
	}
	stateStore := state.NewStore(stateDB, state.StoreOptions{
		DiscardABCIResponses: config.Storage.DiscardABCIResponses,
	})

	return blockStore, stateStore, nil
}

func loadRollkitStateStore(rootDir, dbPath string) (rollkitstore.Store, error) {
	baseKV, err := rollkitstore.NewDefaultKVStore(rootDir, dbPath, "rollkit")
	if err != nil {
		return nil, err
	}

	store := rollkitstore.New(baseKV)
	return store, nil
}

func rollkitStateFromCometBFTState(cometBFTState state.State) (rollkittypes.State, error) {
	return rollkittypes.State{
		Version: cometBFTState.Version,

		ChainID:         cometBFTState.ChainID,
		InitialHeight:   uint64(cometBFTState.InitialHeight),
		LastBlockHeight: uint64(cometBFTState.LastBlockHeight),
		LastBlockID:     cometBFTState.LastBlockID,
		LastBlockTime:   cometBFTState.LastBlockTime,

		DAHeight: 1,

		ConsensusParams:                  cometBFTState.ConsensusParams.ToProto(),
		LastHeightConsensusParamsChanged: uint64(cometBFTState.LastHeightConsensusParamsChanged),

		LastResultsHash: cometBFTState.LastResultsHash,
		AppHash:         cometBFTState.AppHash,

		Validators:                  cometBFTState.Validators,
		NextValidators:              cometBFTState.NextValidators,
		LastValidators:              cometBFTState.LastValidators,
		LastHeightValidatorsChanged: cometBFTState.LastHeightValidatorsChanged,
	}, nil
}

func rollkitBlockFromCometBFTBlock(block *cometbfttypes.Block, commit cometbfttypes.Commit, validatorSet *cometbfttypes.ValidatorSet) rollkittypes.Block {
	rollkitTxs := make([]rollkittypes.Tx, len(block.Data.Txs))
	for i, tx := range block.Data.Txs {
		rollkitTxs[i] = rollkittypes.Tx(tx)
	}
	return rollkittypes.Block{
		SignedHeader: rollkittypes.SignedHeader{
			Header:     rollkitHeaderFromCometBFTHeader(block.Header),
			Commit:     rollkitCommitFromCometBFTCommit(commit),
			Validators: validatorSet,
		},
		Data: rollkittypes.Data{
			Txs: rollkitTxs,
		},
	}
}

func rollkitCommitFromCometBFTCommit(commit cometbfttypes.Commit) rollkittypes.Commit {
	rollkitSigs := make([]rollkittypes.Signature, len(commit.Signatures))
	for i, sig := range commit.Signatures {
		rollkitSigs[i] = sig.Signature
	}

	return rollkittypes.Commit{
		Signatures: rollkitSigs,
	}
}

func rollkitHeaderFromCometBFTHeader(header cometbfttypes.Header) rollkittypes.Header {
	return rollkittypes.Header{
		BaseHeader: rollkittypes.BaseHeader{
			ChainID: header.ChainID,
			Height:  uint64(header.Height),
			Time:    uint64(header.Time.Unix()),
		},
		Version: rollkittypes.Version{
			Block: header.Version.Block,
			App:   header.Version.App,
		},
		LastHeaderHash: header.ToProto().LastCommitHash,
		DataHash:       header.ToProto().DataHash,
		ConsensusHash:  header.ToProto().ConsensusHash,
		AppHash:        header.ToProto().AppHash,

		LastCommitHash:  header.ToProto().LastCommitHash,
		LastResultsHash: header.ToProto().LastResultsHash,
		ValidatorHash:   header.ToProto().ValidatorsHash,

		ProposerAddress: header.ProposerAddress,
	}
}

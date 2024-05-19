package types

// NewGenesisState creates a new GenesisState instance
func NewGenesisState(params Params, sequencer []Sequencer) *GenesisState {
	return &GenesisState{
		Params:     params,
		Sequencers: sequencer,
	}
}

// DefaultGenesisState gets the raw genesis raw message for testing
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Sequencers: []Sequencer{},
	}
}

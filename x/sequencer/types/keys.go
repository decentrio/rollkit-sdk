package types

const (
	// ModuleName is the name of the sequencer module
	ModuleName = "sequencer"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the sequencer module
	RouterKey = ModuleName
)

var (
	// Keys for store sequencer cons address
	SequencerConsAddrKey = []byte{0x11}

	// Keys for store last validator set
	LastValidatorSetKey = []byte{0x12}

	// Keys for store next sequencer change height
	NextSequencerChangeHeight = []byte{0x13}

	ParamsKey = []byte{0x14}
)

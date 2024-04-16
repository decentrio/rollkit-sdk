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
)

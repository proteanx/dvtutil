package dvtutil

import "github.com/martinboehm/btcd/wire"

// Node_BitcoinCash service bit
const SFNodeBitcoinCash wire.ServiceFlag = 1 << 5

const (
	// MainNet represents the main devault network.
	MainnetMagic wire.BitcoinNet = 0xde3a9c03

	// Testnet represents the test network (version 3).
	TestnetMagic wire.BitcoinNet = 0xf4f3e5f4

	// Regtest represents the regression test network.
	Regtestmagic wire.BitcoinNet = 0xfabfb5da
)

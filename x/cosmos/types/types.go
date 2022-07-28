package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Success          = "success"
	GasFailure       = "gas failure"
	SequenceMismatch = "sequence mismatch"
	KeeperFailure    = "keeper failure"
	NotSuccess       = "not success"
)

// UInt64FromBytes create uint from binary big endian representation
func UInt64FromBytes(s []byte) uint64 {
	return binary.BigEndian.Uint64(s)
}

// UInt64Bytes uses the SDK byte marshaling to encode a uint64
func UInt64Bytes(n uint64) []byte {
	return sdk.Uint64ToBigEndian(n)
}

// Int64FromBytes create int64 from binary big endian representation
func Int64FromBytes(s []byte) int64 {
	return int64(binary.BigEndian.Uint64(s))
}

// Int64Bytes uses the SDK byte marshaling to encode a uint64 from int64
func Int64Bytes(n int64) []byte {
	return sdk.Uint64ToBigEndian(uint64(n))
}
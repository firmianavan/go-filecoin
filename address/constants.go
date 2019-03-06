package address

import (
	"encoding/base32"

	"gx/ipfs/QmZp3eKdYQHHAneECmeK6HhiMwTPufmjC8DuuaGKv3unvx/blake2b-simd"
)

var (
	// TODO remove TestAddresses

	// TestAddress is an account with some initial funds in it.
	TestAddress Address
	// TestAddress2 is an account with some initial funds in it.
	TestAddress2 Address

	// NetworkAddress is the filecoin network.
	NetworkAddress Address
	// StorageMarketAddress is the hard-coded address of the filecoin storage market.
	StorageMarketAddress Address
	// PaymentBrokerAddress is the hard-coded address of the filecoin storage market.
	PaymentBrokerAddress Address
)

// PayloadHashLength defines the hash length taken over addresses using the Actor and SECP256K1 protocols.
const PayloadHashLength = 20

// ChecksumHashLength defines the hash length used for calculating address checksums.
const ChecksumHashLength = 4

var payloadHashConfig = &blake2b.Config{Size: PayloadHashLength}
var checksumHashConfig = &blake2b.Config{Size: ChecksumHashLength}

const encodeStd = "abcdefghijklmnopqrstuvwxyz234567"

// AddressEncoding defines the base32 config used for address encoding and decoding.
var AddressEncoding = base32.NewEncoding(encodeStd)

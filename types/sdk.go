package types

import (
	"encoding/hex"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccAddressFromHexUnsafe creates an AccAddress from a HEX-encoded string.
//
// Note, this function is considered unsafe as it may produce an AccAddress from
// otherwise invalid input, such as a transaction hash. Please use
// AccAddressFromBech32.
func AccAddressFromHexUnsafe(address string) (addr sdk.AccAddress, err error) {
	bz, err := addressBytesFromHexString(address)
	return sdk.AccAddress(bz), err
}

func addressBytesFromHexString(address string) ([]byte, error) {
	if len(address) == 0 {
		return nil, errors.New("decoding Bech32 address failed: must provide an address")
	}

	return hex.DecodeString(address)
}

package app

import (
	"github.com/fibonacci-chain/core/app/params"
	enccodec "github.com/fibonacci-chain/core/encoding/codec"
)

// MakeEncodingConfig creates an EncodingConfig for testing.
func MakeEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()

	//support ethsecp256k1
	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	//std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	//std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

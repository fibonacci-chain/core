package bindings

import "github.com/fibonacci-chain/core/x/epoch/types"

type SeiEpochQuery struct {
	// queries the current Epoch
	Epoch *types.QueryEpochRequest `json:"epoch,omitempty"`
}

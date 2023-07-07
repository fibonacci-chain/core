// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/evmos/ethermint/blob/main/LICENSE
package flags

import (
	"github.com/fibonacci-chain/core/server/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// JSON-RPC flags
const (
	JSONRPCEnable              = "json-rpc.enable"
	JSONRPCAPI                 = "json-rpc.api"
	JSONRPCAddress             = "json-rpc.address"
	JSONWsAddress              = "json-rpc.ws-address"
	JSONRPCGasCap              = "json-rpc.gas-cap"
	JSONRPCEVMTimeout          = "json-rpc.evm-timeout"
	JSONRPCTxFeeCap            = "json-rpc.txfee-cap"
	JSONRPCFilterCap           = "json-rpc.filter-cap"
	JSONRPCLogsCap             = "json-rpc.logs-cap"
	JSONRPCBlockRangeCap       = "json-rpc.block-range-cap"
	JSONRPCHTTPTimeout         = "json-rpc.http-timeout"
	JSONRPCHTTPIdleTimeout     = "json-rpc.http-idle-timeout"
	JSONRPCAllowUnprotectedTxs = "json-rpc.allow-unprotected-txs"
	JSONRPCMaxOpenConnections  = "json-rpc.max-open-connections"
	JSONRPCEnableIndexer       = "json-rpc.enable-indexer"
	// JSONRPCEnableMetrics enables EVM RPC metrics server.
	// Set to `metrics` which is hardcoded flag from go-ethereum.
	// https://github.com/ethereum/go-ethereum/blob/master/metrics/metrics.go#L35-L55
	JSONRPCEnableMetrics            = "metrics"
	JSONRPCFixRevertGasRefundHeight = "json-rpc.fix-revert-gas-refund-height"
)

// EVM flags
const (
	EVMTracer         = "evm.tracer"
	EVMMaxTxGasWanted = "evm.max-tx-gas-wanted"
)

func AddJSONRPCFlags(cmd *cobra.Command) (*cobra.Command, error) {
	cmd.PersistentFlags().Bool(JSONRPCEnable, true, "Define if the JSON-RPC server should be enabled")
	cmd.PersistentFlags().StringSlice(JSONRPCAPI, config.GetDefaultAPINamespaces(), "Defines a list of JSON-RPC namespaces that should be enabled")
	cmd.PersistentFlags().String(JSONRPCAddress, config.DefaultJSONRPCAddress, "the JSON-RPC server address to listen on")
	cmd.PersistentFlags().String(JSONWsAddress, config.DefaultJSONRPCWsAddress, "the JSON-RPC WS server address to listen on")
	cmd.PersistentFlags().Uint64(JSONRPCGasCap, config.DefaultGasCap, "Sets a cap on gas that can be used in eth_call/estimateGas unit is aphoton (0=infinite)")     //nolint:lll
	cmd.PersistentFlags().Float64(JSONRPCTxFeeCap, config.DefaultTxFeeCap, "Sets a cap on transaction fee that can be sent via the RPC APIs (1 = default 1 photon)") //nolint:lll
	cmd.PersistentFlags().Int32(JSONRPCFilterCap, config.DefaultFilterCap, "Sets the global cap for total number of filters that can be created")
	cmd.PersistentFlags().Duration(JSONRPCEVMTimeout, config.DefaultEVMTimeout, "Sets a timeout used for eth_call (0=infinite)")
	cmd.PersistentFlags().Duration(JSONRPCHTTPTimeout, config.DefaultHTTPTimeout, "Sets a read/write timeout for json-rpc http server (0=infinite)")
	cmd.PersistentFlags().Duration(JSONRPCHTTPIdleTimeout, config.DefaultHTTPIdleTimeout, "Sets a idle timeout for json-rpc http server (0=infinite)")
	cmd.PersistentFlags().Bool(JSONRPCAllowUnprotectedTxs, config.DefaultAllowUnprotectedTxs, "Allow for unprotected (non EIP155 signed) transactions to be submitted via the node's RPC when the global parameter is disabled") //nolint:lll
	cmd.PersistentFlags().Int32(JSONRPCLogsCap, config.DefaultLogsCap, "Sets the max number of results can be returned from single `eth_getLogs` query")
	cmd.PersistentFlags().Int32(JSONRPCBlockRangeCap, config.DefaultBlockRangeCap, "Sets the max block range allowed for `eth_getLogs` query")
	cmd.PersistentFlags().Int32(JSONRPCMaxOpenConnections, config.DefaultMaxOpenConnections, "Sets the maximum number of simultaneous connections for the server listener") //nolint:lll
	cmd.PersistentFlags().Bool(JSONRPCEnableIndexer, false, "Enable the custom tx indexer for json-rpc")
	cmd.PersistentFlags().Bool(JSONRPCEnableMetrics, false, "Define if EVM rpc metrics server should be enabled")

	if err := viper.BindPFlag(JSONRPCEnable, cmd.PersistentFlags().Lookup(JSONRPCEnable)); err != nil {
		return nil, err
	}

	return cmd, nil
}

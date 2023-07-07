package wasm

import (
	tokenfactorykeeper "github.com/fibonacci-chain/core/x/tokenfactory/keeper"
)

type TokenFactoryWasmQueryHandler struct {
	tokenfactoryKeeper tokenfactorykeeper.Keeper
}

func NewTokenFactoryWasmQueryHandler(keeper *tokenfactorykeeper.Keeper) *TokenFactoryWasmQueryHandler {
	return &TokenFactoryWasmQueryHandler{
		tokenfactoryKeeper: *keeper,
	}
}

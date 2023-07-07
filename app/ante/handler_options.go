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
package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/fibonacci-chain/core/app/antedecorators/depdecorators"

	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"

	evmtypes "github.com/fibonacci-chain/core/x/evm/types"
)

// HandlerOptions2 extend the SDK's AnteHandler options by requiring the IBC
// channel keeper, EVM Keeper and Fee Market Keeper.
type HandlerOptions struct {
	AccountKeeper   evmtypes.AccountKeeper
	BankKeeper      evmtypes.BankKeeper
	IBCKeeper       *ibckeeper.Keeper
	FeeMarketKeeper FeeMarketKeeper
	ParamsKeeper    paramskeeper.Keeper

	EvmKeeper              EVMKeeper
	FeegrantKeeper         ante.FeegrantKeeper
	SignModeHandler        authsigning.SignModeHandler
	SigGasConsumer         func(meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params) error
	MaxTxGasWanted         uint64
	ExtensionOptionChecker ExtensionOptionChecker
	TxFeeChecker           ante.TxFeeChecker
	DisabledAuthzMsgs      []string
}

func (options HandlerOptions) validate() error {
	if options.AccountKeeper == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.FeeMarketKeeper == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "fee market keeper is required for AnteHandler")
	}
	if options.EvmKeeper == nil {
		return errorsmod.Wrap(errortypes.ErrLogic, "evm keeper is required for AnteHandler")
	}
	return nil
}

func newEthAnteHandler(options HandlerOptions) sdk.AnteHandler {
	a, _ := sdk.ChainAnteDecorators(
		sdk.DefaultWrappedAnteDecorator(NewEthSetUpContextDecorator(options.EvmKeeper)),                         // outermost AnteDecorator. SetUpContext must be called first
		sdk.DefaultWrappedAnteDecorator(NewEthMempoolFeeDecorator(options.EvmKeeper)),                           // Check eth effective gas price against minimal-gas-prices
		sdk.DefaultWrappedAnteDecorator(NewEthMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper)), // Check eth effective gas price against the global MinGasPrice
		sdk.DefaultWrappedAnteDecorator(NewEthValidateBasicDecorator(options.EvmKeeper)),
		sdk.DefaultWrappedAnteDecorator(NewEthSigVerificationDecorator(options.EvmKeeper)),
		sdk.DefaultWrappedAnteDecorator(NewEthAccountVerificationDecorator(options.AccountKeeper, options.EvmKeeper)),
		sdk.DefaultWrappedAnteDecorator(NewCanTransferDecorator(options.EvmKeeper)),
		sdk.DefaultWrappedAnteDecorator(NewEthGasConsumeDecorator(options.EvmKeeper, options.MaxTxGasWanted)),
		sdk.DefaultWrappedAnteDecorator(NewEthIncrementSenderSequenceDecorator(options.AccountKeeper)), // innermost AnteDecorator.
		sdk.DefaultWrappedAnteDecorator(NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper)),
		sdk.DefaultWrappedAnteDecorator(NewEthEmitEventDecorator(options.EvmKeeper)), // emit eth tx hash and index at the very last ante handler.
	)
	return a
}

func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	a, _ := sdk.ChainAnteDecorators(
		sdk.DefaultWrappedAnteDecorator(RejectMessagesDecorator{}), // reject MsgEthereumTxs
		// disable the Msg types that cannot be included on an authz.MsgExec msgs field
		sdk.DefaultWrappedAnteDecorator(NewAuthzLimiterDecorator(options.DisabledAuthzMsgs)),
		sdk.CustomDepWrappedAnteDecorator(authante.NewDefaultSetUpContextDecorator(), depdecorators.GasMeterSetterDecorator{}), // outermost AnteDecorator. SetUpContext must be called first
		sdk.DefaultWrappedAnteDecorator(NewExtensionOptionsDecorator(options.ExtensionOptionChecker)),
		sdk.DefaultWrappedAnteDecorator(ante.NewValidateBasicDecorator()),
		sdk.DefaultWrappedAnteDecorator(ante.NewTxTimeoutHeightDecorator()),
		sdk.DefaultWrappedAnteDecorator(NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper)),
		sdk.DefaultWrappedAnteDecorator(ante.NewValidateMemoDecorator(options.AccountKeeper)),
		sdk.DefaultWrappedAnteDecorator(ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper)),
		sdk.DefaultWrappedAnteDecorator(ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper, options.ParamsKeeper, options.TxFeeChecker)),
		// SetPubKeyDecorator must be called before all signature verification decorators
		sdk.DefaultWrappedAnteDecorator(ante.NewSetPubKeyDecorator(options.AccountKeeper)),
		sdk.DefaultWrappedAnteDecorator(ante.NewValidateSigCountDecorator(options.AccountKeeper)),
		sdk.DefaultWrappedAnteDecorator(ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer)),
		sdk.DefaultWrappedAnteDecorator(ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler)),
		sdk.DefaultWrappedAnteDecorator(ante.NewIncrementSequenceDecorator(options.AccountKeeper)),
		//sdk.DefaultWrappedAnteDecorator(ibcante.NewRedundantRelayDecorator(options.IBCKeeper)),
		sdk.DefaultWrappedAnteDecorator(NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper)),
	)
	return a
}

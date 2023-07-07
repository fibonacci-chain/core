package main

import (
	"os"

	"github.com/fibonacci-chain/core/app/params"
	"github.com/fibonacci-chain/core/cmd/fbchaind/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/fibonacci-chain/core/app"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

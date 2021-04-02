package main

import (
	"os"

	"github.com/compamint/jetcoin/cmd/jetcoind/cmd"
    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/compamint/jetcoin/app"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
    if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

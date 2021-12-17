package main

import (
	"go-blockchain/core"
	"os"
)

func main() {
	defer os.Exit(0)
	chain := core.InitBlockchain()
	defer chain.Database.Close()
	cli := core.CommandLine{chain}
	cli.Run()
}

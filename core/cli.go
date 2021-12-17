package core

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type CommandLine struct {
	Blockchain *Blockchain
}

func (cli *CommandLine) printUsage() {
	fmt.Print("Usage:")
	fmt.Println("  add -block BLOCK_DATA -> add a block to the chain")
	fmt.Println("  print -> prints the block in the chain")
}

func (cli *CommandLine) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.Blockchain.AddBlock(data)
	fmt.Println("Successfully add a new block to the blockchain")
}

func (cli *CommandLine) printChain() {
	iter := cli.Blockchain.Iterator()
	fmt.Println("printing iter last hash")
	// fmt.Printf("%x\n", iter.CurrentHash)
	for {
		
		block := iter.Next()
		fmt.Println("==================================")
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		// proving block
		pow := NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) Run() {
	cli.ValidateArgs()

	addBlockCmd 	:= flag.NewFlagSet("add"	, flag.ExitOnError)
	printChainCmd 	:= flag.NewFlagSet("print"	, flag.ExitOnError)
	
	addBlockData := addBlockCmd.String("block", "", "Block Data")

	switch os.Args[1] {

	case "add": 
		err := addBlockCmd.Parse(os.Args[2:])
		Handle(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()

	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
 	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
	
}
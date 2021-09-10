package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) printUsage() {
	fmt.Println("使用信息：")
	fmt.Println("addblock -data 交易信息-添加区块")
	fmt.Println("printchain -遍历区块并打印")
}
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}
func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("成功添加新区块")
}
func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("上一个区块的Hash：%x\n", block.PreBlockHash)
		fmt.Printf("区块信息：%s\n", block.Data)
		fmt.Printf("当前区块的Hash：%x\n", block.Hash)
		pow := NewProofOfWOrk(block)
		fmt.Println("Pow:", pow.Validate())
		fmt.Println()
		if len(block.PreBlockHash) == 0 {
			break
		}
	}
}
func (cli *CLI) Run() {
	cli.validateArgs()
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "区块信息")
	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		cli.printChain()
	default:
		cli.printUsage()
		os.Exit(1)
	}
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}
}

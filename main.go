package main

func main() {
	blockchain := NewBlockChain()
	blockchain.AddBlock("a")
	blockchain.AddBlock("b")
	blockchain.AddBlock("c")

	// for _, block := range blockchain.blocks {
	// 	fmt.Println(block.PreBlockHash)
	// 	fmt.Println(block.Data)
	// 	fmt.Println(block.Hash)
	// 	pow := NewProofOfWOrk(block)
	// 	fmt.Println(strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }

}

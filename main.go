package main

import (
	"fmt"
	"naivecoin-go/block"
)

func main() {
	bc := block.BlockChain{}
	bc.Init("first block")
	latestBlock := bc.GetLatestBlock()
	fmt.Println(block.CalculateHash(latestBlock.Index, latestBlock.PreviousHash, latestBlock.Timestamp, latestBlock.Data))
}

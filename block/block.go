package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Block struct {
	Index        int64
	Hash         string
	PreviousHash string
	Timestamp    int64
	Data         string
}

func CalculateHash(index int64, previousHash string, timestamp int64, data string) string {
	sum := []byte(strconv.Itoa(int(index)))
	sum = append(sum, []byte(previousHash)...)
	sum = append(sum, []byte(strconv.Itoa(int(timestamp)))...)
	sum = append(sum, []byte(data)...)
	res := sha256.Sum256(sum)
	return hex.EncodeToString(res[:])
}

type BlockChain struct {
	chain []Block
}

func (bc *BlockChain) Init(data string) *BlockChain {
	first := Block{
		Index:        0,
		Hash:         "0",
		PreviousHash: "0",
		Timestamp:    time.Now().Unix(),
		Data:         data,
	}
	first.Hash = CalculateHash(first.Index, first.PreviousHash, first.Timestamp, first.Data)
	bc.chain = append(bc.chain, first)
	return bc
}

func (bc BlockChain) GetLatestBlock() Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc BlockChain) GenerateNextBlock(blockData string) Block {
	previousBlock := bc.GetLatestBlock()
	nextIndex := previousBlock.Index + 1
	nextTimestamp := time.Now().Unix()
	nextHash := CalculateHash(nextIndex, previousBlock.PreviousHash, nextTimestamp, blockData)
	newBlock := Block{
		Index:        nextIndex,
		Hash:         nextHash,
		PreviousHash: previousBlock.Hash,
		Timestamp:    nextTimestamp,
		Data:         blockData,
	}
	return newBlock
}

func (bc *BlockChain) ReplaceChain(newBlockChain []Block) {
	if IsValidChain(newBlockChain) && len(newBlockChain) > len(bc.chain) {
		fmt.Println("Received blockchain is valid. Replacing current blockchain with received blockchain")
		bc.chain = newBlockChain
		// broadcast
	} else {
		fmt.Println("Received blockchain invalid")
	}
}

func IsValidNewBlock(newBlock Block, previousBlock Block) bool {
	if previousBlock.Index+1 != newBlock.Index {
		fmt.Println("invalid index")
		return false
	} else if previousBlock.Hash != newBlock.PreviousHash {
		fmt.Println("invalid previousHash")
		return false
	} else if CalculateHash(newBlock.Index, newBlock.PreviousHash, newBlock.Timestamp, newBlock.Data) != newBlock.Hash {
		fmt.Println("invalid hash")
		return false
	}
	return true
}

func IsValidBlockStruct(block Block) bool {
	return reflect.TypeOf(block.Index).Kind() == reflect.Int64 &&
		reflect.TypeOf(block.Hash).Kind() == reflect.String &&
		reflect.TypeOf(block.PreviousHash).Kind() == reflect.String &&
		reflect.TypeOf(block.Timestamp).Kind() == reflect.Int64 &&
		reflect.TypeOf(block.Data).Kind() == reflect.String
}

func IsValidChain(chain []Block) bool {
	for i := 0; i < len(chain); i++ {
		if !IsValidNewBlock(chain[i], chain[i-1]) {
			return false
		}
	}
	return true
}

var Bc = BlockChain{}

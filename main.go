package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/orm"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/parser"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	nodeURL        = ""
	checkpointFile = "checkpoint.txt"
)

var (
	ContractABI *abi.ABI

	EventSig common.Hash

	ChainID = 11155111

	ContractAddress = "0xB74D5Dba3081bCaDb5D4e1CC77Cc4807E1c4ecf8"
)

func main() {

	fromBlock := big.NewInt(2748098)
	toBlock := big.NewInt(2800000)
	blockStep := big.NewInt(300)
	db, err := orm.InitDB()
	if err != nil {
		log.Fatal("failed to init db", "err", err)
	}
	err = db.Table("raw_events").AutoMigrate(&orm.RawEvents{})
	if err != nil {
		log.Fatal("failed to AutoMigrate db", "err", err)
	}
	defer orm.CloseDB(db)
	rawEventsOrm := orm.NewRawEvents(db)
	if err != nil {
		log.Fatal("failed to init rawEventsOrm", "err", err)
	}
	ProcessBlocks(fromBlock, toBlock, blockStep, rawEventsOrm)

}

// ProcessBlocks processes blocks in the specified range
func ProcessBlocks(fromBlock, toBlock, blockStep *big.Int, rawEventsOrm *orm.RawEvents) {
	for currentBlock := new(big.Int).Set(fromBlock); currentBlock.Cmp(toBlock) <= 0; {
		endBlock := new(big.Int).Add(currentBlock, blockStep)
		if endBlock.Cmp(toBlock) > 0 {
			endBlock = toBlock
		}
		if endBlock.Cmp(currentBlock) <= 0 {
			break
		}
		err := parser.FetchAndProcessEvents(nodeURL, ContractAddress, new(big.Int).Add(currentBlock, big.NewInt(1)), endBlock, rawEventsOrm)
		if err != nil {
			fmt.Println("currentBlock", new(big.Int).Add(currentBlock, big.NewInt(1)))
			fmt.Println("endBlock", endBlock)
			log.Fatalf("Failed to subscribe and process events: %v", err)
		}
		fmt.Printf("Processed blocks from %s to %s\n", currentBlock.String(), endBlock.String())

		time.Sleep(2 * time.Second)
		currentBlock = endBlock
	}
}

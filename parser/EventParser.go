package parser

import (
	"context"
	"fmt"
	"math/big"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/orm"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SubscribeAndProcessEvents subscribes to the specified contract's events and processes them.
func FetchAndProcessEvents(nodeURL, contractAddress string, fromBlock, toBlock *big.Int, RawEventsOrm *orm.RawEvents) error {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	address := common.HexToAddress(contractAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
		FromBlock: fromBlock,
		ToBlock:   toBlock,
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to filter logs: %v", err)
	}
	Events := make([]*orm.RawEvents, 0)
	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		// filter out the events we are interested in
	}
	err = RawEventsOrm.InsertRawEvents(context.Background(), "raw_events", Events)
	if err != nil {
		return fmt.Errorf("failed to insert bridge events: %v", err)
	}
	return nil
}

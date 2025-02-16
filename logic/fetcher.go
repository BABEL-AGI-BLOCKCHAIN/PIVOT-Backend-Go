package logic

import (
	"context"
	"math/big"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/common/config"
	backendabi "github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/contract/abi"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/orm"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const ReorgSafeDepth = 64

// FilterResultL1 fetcher result
type FilterResult struct {
	createTopicEvents        []*orm.RawEvent
	investEvents             []*orm.RawEvent
	withdrawEvents           []*orm.RawEvent
	withdrawCommissionEvents []*orm.RawEvent
}

// FetcherLogic the L1 fetcher logic
type FetcherLogic struct {
	cfg          *config.ListenerConfig
	client       *ethclient.Client
	addressList  []common.Address
	parser       *EventParser
	db           *gorm.DB
	rawEventsOrm *orm.RawEvent
}

func NewFetcherLogic(cfg *config.ListenerConfig, db *gorm.DB, client *ethclient.Client) *FetcherLogic {
	addressList := []common.Address{
		common.HexToAddress(cfg.TopicContractAddr),
	}
	f := &FetcherLogic{
		db:           db,
		rawEventsOrm: orm.NewRawEvents(db),
		cfg:          cfg,
		client:       client,
		addressList:  addressList,
		parser:       NewEventParser(cfg),
	}

	return f
}
func (f *FetcherLogic) Fetcher(ctx context.Context, from, to uint64, lastBlockHash common.Hash) (bool, uint64, common.Hash, *FilterResult, error) {
	logrus.Info("fetch and save events from: ", from, "to: ", to)

	isReorg, reorgHeight, blockHash, _, getErr := f.getBlocksAndDetectReorg(ctx, from, to, lastBlockHash)
	if getErr != nil {
		log.Error("L1Fetcher getBlocksAndDetectReorg failed", "from", from, "to", to, "error", getErr)
		return false, 0, common.Hash{}, nil, getErr
	}

	if isReorg {
		return isReorg, reorgHeight, blockHash, nil, nil
	}

	eventLogs, err := f.fetcherLogs(ctx, from, to)

	if err != nil {
		log.Error("L1Fetcher l1FetcherLogs failed", "from", from, "to", to, "error", err)
		return false, 0, common.Hash{}, nil, err
	}

	createTopicEvents, investEvents, withdrawEvets, withdrawCommissionEvents, err := f.parser.ParseRawEventLogs(ctx, eventLogs)

	if err != nil {
		log.Error("failed to parse L1 cross chain event logs", "from", from, "to", to, "err", err)
		return false, 0, common.Hash{}, nil, err
	}

	res := FilterResult{
		createTopicEvents:        createTopicEvents,
		investEvents:             investEvents,
		withdrawEvents:           withdrawEvets,
		withdrawCommissionEvents: withdrawCommissionEvents,
	}

	return false, 0, blockHash, &res, nil
}
func (f *FetcherLogic) getBlocksAndDetectReorg(ctx context.Context, from, to uint64, lastBlockHash common.Hash) (bool, uint64, common.Hash, []*ethtypes.Block, error) {
	blocks, err := utils.GetBlocksInRange(ctx, f.client, from, to)
	if err != nil {
		log.Error("failed to get blocks in range ", " from ", from, " to ", to, " err ", err)
		return false, 0, common.Hash{}, nil, err
	}

	for _, block := range blocks {
		if block.ParentHash() != lastBlockHash {
			log.Warn(" reorg detected", "reorg height", block.NumberU64()-1, "expected hash", block.ParentHash().String(), "local hash", lastBlockHash.String())
			var resyncHeight uint64
			if block.NumberU64() > ReorgSafeDepth+1 {
				resyncHeight = block.NumberU64() - ReorgSafeDepth - 1
			}
			header, err := f.client.HeaderByNumber(ctx, new(big.Int).SetUint64(resyncHeight))
			if err != nil {
				log.Error("failed to get header by number", "block number", resyncHeight, "err", err)
				return false, 0, common.Hash{}, nil, err
			}
			return true, resyncHeight, header.Hash(), nil, nil
		}
		lastBlockHash = block.Hash()
	}

	return false, 0, lastBlockHash, blocks, nil
}

func (f *FetcherLogic) fetcherLogs(ctx context.Context, from, to uint64) ([]ethtypes.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(from), // inclusive
		ToBlock:   new(big.Int).SetUint64(to),   // inclusive
		Addresses: f.addressList,
		Topics:    make([][]common.Hash, 1),
	}
	logrus.Info("fetcherLogs ", " from ", from, " to ", to, "on address ", f.addressList)
	query.Topics[0] = make([]common.Hash, 4)
	query.Topics[0][0] = backendabi.CreateTopicEventSig
	query.Topics[0][1] = backendabi.InvestEventSig
	query.Topics[0][2] = backendabi.WithdrawEventSig
	query.Topics[0][3] = backendabi.WithdrawCommissionEventSig

	eventLogs, err := f.client.FilterLogs(ctx, query)
	if err != nil {
		log.Error("failed to filter event logs", "from", from, "to", to, "err", err)
		return nil, err
	}
	return eventLogs, nil
}

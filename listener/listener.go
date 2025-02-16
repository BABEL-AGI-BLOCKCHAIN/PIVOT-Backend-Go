package listener

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/common/config"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/logic"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/orm"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/utils"
	"github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
)

/*
ERC20
0x83F3c5020Ef0f44C8Ef4993124740D3fe8D1470C

SBT
0x9b11f74888dF35573B934567088578eEe485B663

topic
0x9Af4f4b7C831b0c79574CCDE7C04e33F99BF6438
*/
const (
	// EventTypeCreateTopic represents the event type for creating a topic.
	ReorgSafeDepth = 64
)

type Listener struct {
	ctx    context.Context
	cfg    *config.ListenerConfig
	client *ethclient.Client

	fetcherLogic      *logic.FetcherLogic
	syncHeight        uint64
	lastSyncBlockHash common.Hash

	eventUpdateLogic *logic.EventUpdateLogic
	rawEventsOrm     *orm.RawEvent
	mu               sync.Mutex // Add a mutex to ensure sequential execution

}

func NewListener(ctx context.Context, cfg *config.ListenerConfig, ethClient *ethclient.Client, db *gorm.DB) (*Listener, error) {
	c := &Listener{
		ctx:              ctx,
		cfg:              cfg,
		client:           ethClient,
		fetcherLogic:     logic.NewFetcherLogic(cfg, db, ethClient),
		rawEventsOrm:     orm.NewRawEvents(db),
		eventUpdateLogic: logic.NewEventUpdateLogic(db),
	}
	return c, nil
}

func (c *Listener) Start() {
	messageSyncedHeight, dbErr := c.eventUpdateLogic.GetSyncHeight(c.ctx)
	if dbErr != nil {
		logrus.Fatal("failed to get sync height", "err", dbErr)
	}
	syncHeight := messageSyncedHeight
	if syncHeight < ReorgSafeDepth {
		syncHeight = 0
	} else {
		syncHeight -= ReorgSafeDepth
	}
	if c.cfg.StartHeight > syncHeight {
		syncHeight = c.cfg.StartHeight - 1
	}

	// Sync from an older block to prevent reorg during restart.

	header, err := c.client.HeaderByNumber(c.ctx, new(big.Int).SetUint64(syncHeight))
	if err != nil {
		log.Crit("failed to get header by number", "block number", syncHeight, "err", err)
		return
	}

	c.updateSyncHeight(syncHeight, header.Hash())

	log.Info("Start message fetcher",
		"message synced height", messageSyncedHeight,
		"config start height", c.cfg.StartHeight,
		"sync start height", c.syncHeight+1,
	)

	tick := time.NewTicker(time.Duration(c.cfg.BlockTime) * time.Second)
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				tick.Stop()
				return
			case <-tick.C:
				c.mu.Lock() // Lock the mutex before starting the task
				c.fetchAndSaveEvents(c.cfg.Confirmation)
				c.mu.Unlock() // Unlock the mutex after the task is done
			}
		}
	}()
}
func (w *Listener) updateSyncHeight(height uint64, blockHash common.Hash) {
	w.lastSyncBlockHash = blockHash
	w.syncHeight = height
}

func (w *Listener) ChainID(ctx context.Context) (*big.Int, error) {
	return w.client.ChainID(ctx)
}

func (w *Listener) Close() {
	w.client.Close()
}
func (c *Listener) fetchAndSaveEvents(confirmation uint64) {
	startHeight := c.syncHeight + 1
	endHeight, rpcErr := utils.GetBlockNumber(c.ctx, c.client, confirmation)
	if rpcErr != nil {
		logrus.Error("failed to get block number", " confirmation ", confirmation, " err ", rpcErr)
		return
	}

	logrus.Info("fetch and save missing events, ", "start height: ", startHeight, " , end height: ", endHeight, " confirmation: ", confirmation)

	for from := startHeight; from <= endHeight; from += c.cfg.FetchLimit {
		to := from + c.cfg.FetchLimit - 1
		if to > endHeight {
			to = endHeight
		}
		time.Sleep(2 * time.Second)

		isReorg, resyncHeight, lastBlockHash, l1FetcherResult, fetcherErr := c.fetcherLogic.Fetcher(c.ctx, from, to, c.lastSyncBlockHash)
		if fetcherErr != nil {
			logrus.Error("failed to fetch events ", "from ", from, " to ", to, " err ", fetcherErr)
			return
		}

		if isReorg {
			logrus.Warn("reorg happened, resync  events: ", "from", from, " to", to, " resync height ", resyncHeight, " last block hash ", lastBlockHash)
			c.updateSyncHeight(resyncHeight, lastBlockHash)
			return
		}

		if insertUpdateErr := c.eventUpdateLogic.InsertRawEvents(c.ctx, l1FetcherResult); insertUpdateErr != nil {
			logrus.Error("failed to save events ", "from ", from, " to ", to, " err ", insertUpdateErr)
			return
		}

		c.updateSyncHeight(to, lastBlockHash)
	}
}

package listener

import (
	"context"
	"fmt"
	"math/big"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/contract"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/orm"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/parser"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
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
	contractAddress = "0x9Af4f4b7C831b0c79574CCDE7C04e33F99BF6438"
)

type Listener struct {
	ctx context.Context
	//cfg         *evm.GethConfig
	Client       *ethclient.Client
	Parser       *parser.Parser
	rawEventsOrm *orm.RawEvent
}

func NewListener(ctx context.Context, ethClient *ethclient.Client, db *gorm.DB) (*Listener, error) {

	c := &Listener{
		ctx: ctx,
		//cfg:                cfg,
		Client:       ethClient,
		Parser:       parser.NewParser(),
		rawEventsOrm: orm.NewRawEvents(db),
	}
	return c, nil
}

func (w *Listener) Run(ctx context.Context) error {
	createTopicChan := make(chan *contract.PivotTopicCreateTopic)
	investChan := make(chan *contract.PivotTopicInvest)
	withdrawChan := make(chan *contract.PivotTopicWithdraw)
	withdrawCommissionChan := make(chan *contract.PivotTopicWithdrawCommission)

	if w.Client.Client().SupportsSubscriptions() {
		filterer, err := contract.NewPivotTopicFilterer(common.HexToAddress(contractAddress), w.Client)
		if err != nil {
			return nil
		}

		createTopicSub, err := filterer.WatchCreateTopic(&bind.WatchOpts{Context: ctx}, createTopicChan, nil)
		if err != nil {
			return err
		}
		investSub, err := filterer.WatchInvest(&bind.WatchOpts{Context: ctx}, investChan, nil, nil)
		if err != nil {
			return err
		}
		withdrawSub, err := filterer.WatchWithdraw(&bind.WatchOpts{Context: ctx}, withdrawChan, nil)
		if err != nil {
			return err
		}
		withdrawCommissionSub, err := filterer.WatchWithdrawCommission(&bind.WatchOpts{Context: ctx}, withdrawCommissionChan, nil)
		if err != nil {
			return err
		}

		go func() {
			defer createTopicSub.Unsubscribe()
			defer investSub.Unsubscribe()
			defer withdrawSub.Unsubscribe()
			defer withdrawCommissionSub.Unsubscribe()
			for {
				fmt.Println("Listening for events...")
				select {
				case msg := <-createTopicChan:
					fmt.Println("createTopicChan")
					if msg == nil {
						continue
					}
					w.handleCreateTopicMessage(msg)

				case msg := <-investChan:
					if msg == nil {
						continue
					}
					w.handleInvestMessage(msg)

				case msg := <-withdrawChan:
					if msg == nil {
						continue
					}
					w.handleWithdrawMessage(msg)

				case msg := <-withdrawCommissionChan:
					if msg == nil {
						continue
					}
					w.handleWithdrawCommissionMessage(msg)
				case subErr := <-createTopicSub.Err():
					filterer, err := contract.NewPivotTopicFilterer(common.HexToAddress(contractAddress), w.Client)
					logrus.Errorf("L1 createTopic subscription failed: %v, Resubscribing...", subErr)
					createTopicSub, err = filterer.WatchCreateTopic(&bind.WatchOpts{Context: ctx}, createTopicChan, nil)
					if err != nil {
						logrus.Errorf("Resubscribe failed: %v", err)
						return
					}
				case subErr := <-investSub.Err():
					filterer, err := contract.NewPivotTopicFilterer(common.HexToAddress(contractAddress), w.Client)
					logrus.Errorf("L1 invest subscription failed: %v, Resubscribing...", subErr)
					investSub, err = filterer.WatchInvest(&bind.WatchOpts{Context: ctx}, investChan, nil, nil)
					if err != nil {
						logrus.Errorf("Resubscribe failed: %v", err)
						return
					}
				case subErr := <-withdrawSub.Err():
					filterer, err := contract.NewPivotTopicFilterer(common.HexToAddress(contractAddress), w.Client)
					logrus.Errorf("L1 withdraw subscription failed: %v, Resubscribing...", subErr)
					withdrawSub, err = filterer.WatchWithdraw(&bind.WatchOpts{Context: ctx}, withdrawChan, nil)
					if err != nil {
						logrus.Errorf("Resubscribe failed: %v", err)
						return
					}
				case subErr := <-withdrawCommissionSub.Err():
					filterer, err := contract.NewPivotTopicFilterer(common.HexToAddress(contractAddress), w.Client)
					logrus.Errorf("L1 withdrawCommission subscription failed: %v, Resubscribing...", subErr)
					withdrawCommissionSub, err = filterer.WatchWithdrawCommission(&bind.WatchOpts{Context: ctx}, withdrawCommissionChan, nil)
					if err != nil {
						logrus.Errorf("Resubscribe failed: %v", err)
						return
					}

				case <-ctx.Done():
					return
				}
			}
		}()
	}
	return nil
}

func (w *Listener) handleCreateTopicMessage(
	msg *contract.PivotTopicCreateTopic,
) error {
	fmt.Println("handleCreateTopicMessage")
	bridgeEvents, err := w.Parser.ParseCreateTopicEventToRawBridgeEvents(w.ctx, msg)
	if err != nil {
		logrus.Errorf("Failed to parse downward message: %v", err)
	}
	err = w.rawEventsOrm.InsertRawEvents(w.ctx, "raw_events", bridgeEvents)
	if err != nil {
		return err
	}
	return nil
}
func (w *Listener) handleInvestMessage(
	msg *contract.PivotTopicInvest,
) error {
	bridgeEvents, err := w.Parser.ParseInvestEventToRawBridgeEvents(w.ctx, msg)
	if err != nil {
		logrus.Errorf("Failed to parse invest message: %v", err)
	}
	err = w.rawEventsOrm.InsertRawEvents(w.ctx, "raw_events", bridgeEvents)
	if err != nil {
		return err
	}
	return nil
}

func (w *Listener) handleWithdrawMessage(
	msg *contract.PivotTopicWithdraw,
) error {
	bridgeEvents, err := w.Parser.ParseWithdrawEventToRawBridgeEvents(w.ctx, msg)
	if err != nil {
		logrus.Errorf("Failed to parse withdraw message: %v", err)
	}
	err = w.rawEventsOrm.InsertRawEvents(w.ctx, "raw_events", bridgeEvents)
	if err != nil {
		return err
	}
	return nil
}

func (w *Listener) handleWithdrawCommissionMessage(
	msg *contract.PivotTopicWithdrawCommission,
) error {
	bridgeEvents, err := w.Parser.ParseWithdrawCommissionEventToRawBridgeEvents(w.ctx, msg)
	if err != nil {
		logrus.Errorf("Failed to parse withdraw commission message: %v", err)
	}
	err = w.rawEventsOrm.InsertRawEvents(w.ctx, "raw_events", bridgeEvents)
	if err != nil {
		return err
	}
	return nil
}
func (w *Listener) ChainID(ctx context.Context) (*big.Int, error) {
	return w.Client.ChainID(ctx)
}

func (w *Listener) Close() {
	w.Client.Close()
}

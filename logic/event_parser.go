package logic

import (
	"context"
	"time"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/common/config"
	backendabi "github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/contract/abi"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/orm"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/types"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/utils"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/sirupsen/logrus"
)

/*
   event CreateTopic(
       address indexed promoter,
       uint256 topicId,
       uint256 investment,
       uint256 position,
       uint256 nonce
   );
   event Invest(
       address indexed investor,
       uint256 indexed topicId,
       uint256 amount,
       uint256 position,
       uint256 nonce
   );
   event Withdraw(
		address indexed to,
		uint256 amount,
		uint256 nonce
	 );

   event WithdrawCommission(
       address indexed owner,
       uint256 amount,
       uint256 nonce
   );

type RawEvent struct {
	db                *gorm.DB   `gorm:"column:-"`
	ID                uint64     `json:"id" gorm:"column:id;primary_key;autoIncrement"` // primary key in the database
	EventType         int        `json:"event_type" gorm:"column:event_type"`           //
	ChainID           int        `json:"chain_id" gorm:"column:chain_id"`               //
	ContractAddress   string     `json:"contract_address" gorm:"column:contract_address"`
	TokenAddress      string     `json:"token_address" gorm:"column:token_address"`
	TxHash            string     `json:"tx_hash" gorm:"column:tx_hash"`
	BlockNumber       uint64     `json:"block_number" gorm:"column:block_number"`
	MsgValue          string     `json:"msg_value" gorm:"column:msg_value"`
	Timestamp         uint64     `json:"timestamp" gorm:"column:timestamp"`
	Sender            string     `json:"sender" gorm:"column:sender"` // sender address
	MessageTopicId    uint64     `json:"message_topic_id" gorm:"column:message_topic_id"`
	MessageSender     string     `json:"message_promoter" gorm:"column:message_sender"`
	MessagePosition   int        `json:"message_position" gorm:"column:message_position"`
	MessageNonce      int        `json:"message_nonce" gorm:"column:message_nonce"`
	MessageWithdrawTo string     `json:"message_withdrawto" gorm:"column:message_withdrawto"`
	MessageAmount     int        `json:"message_withdraw_amount" gorm:"column:message_amount"`
	CreatedAt         time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt         *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Remark            string     `json:"remark" gorm:"column:remark"`
	ProcessStatus     int        `json:"process_status" gorm:"column:process_status"`
}
*/

type EventParser struct {
	cfg *config.ListenerConfig
}

func NewEventParser(cfg *config.ListenerConfig) *EventParser {
	return &EventParser{
		cfg: cfg,
	}
}

// ParseL1CrossChainEventLogs parse l1 cross chain event logs
func (e *EventParser) ParseRawEventLogs(ctx context.Context, logs []ethtypes.Log) ([]*orm.RawEvent, []*orm.RawEvent, []*orm.RawEvent, []*orm.RawEvent, error) {
	createTopicEvents, investEvents, withdrawEvets, withdrawCommissionEvents, err := e.ParseSingleRawBridgeEventLogs(ctx, logs)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return createTopicEvents, investEvents, withdrawEvets, withdrawCommissionEvents, nil
}
func (e *EventParser) ParseSingleRawBridgeEventLogs(ctx context.Context, logs []ethtypes.Log) ([]*orm.RawEvent, []*orm.RawEvent, []*orm.RawEvent, []*orm.RawEvent, error) {
	var createTopicEvents []*orm.RawEvent
	var investEvents []*orm.RawEvent
	var withdrawEvents []*orm.RawEvent
	var withdrawCommissionEvents []*orm.RawEvent

	for _, vlog := range logs {
		switch vlog.Topics[0] {
		case backendabi.CreateTopicEventSig:
			logrus.Info("Found CreateTopicEventSig")
			event := backendabi.PivotTopicCreateTopic{}

			if err := utils.UnpackLog(backendabi.IPivotTopicABI, &event, "CreateTopic", vlog); err != nil {
				logrus.Error("Failed to unpack CreateTopic event", "err", err)
				return nil, nil, nil, nil, err
			}
			logrus.Info("CreateTopicEventSig ", " event ", vlog.TxHash)
			rawEvent := &orm.RawEvent{
				EventType:       int(types.CreateTopic),
				ChainID:         types.Sepolia,
				ContractAddress: types.ContractAddress,
				TxHash:          vlog.TxHash.String(),
				Timestamp:       uint64(time.Now().Unix()),
				BlockNumber:     vlog.BlockNumber,
				MessageTopicId:  event.TopicId.Uint64(),
				Sender:          event.Promoter.String(),
				MessageSender:   event.Promoter.String(),
				MessageAmount:   int(event.Investment.Int64()),
				MessagePosition: int(event.Position.Int64()),
				MessageNonce:    int(event.Nonce.Int64()),
				CreatedAt:       time.Now().UTC(),
				UpdatedAt:       time.Now().UTC(),
				ProcessStatus:   int(types.UnProcessed),
			}
			createTopicEvents = append(createTopicEvents, rawEvent)

		case backendabi.InvestEventSig:
			logrus.Info("Found CreateTopicEventSig")
			event := backendabi.PivotTopicInvest{}
			if err := utils.UnpackLog(backendabi.IPivotTopicABI, &event, "Invest", vlog); err != nil {
				log.Error("Failed to unpack Invest event", "err", err)
				return nil, nil, nil, nil, err
			}

			rawEvent := &orm.RawEvent{
				EventType:       int(types.Invest),
				ChainID:         types.Sepolia,
				ContractAddress: types.ContractAddress,
				TxHash:          vlog.TxHash.String(),
				Timestamp:       uint64(time.Now().Unix()),
				BlockNumber:     vlog.BlockNumber,
				MessageTopicId:  event.TopicId.Uint64(),
				Sender:          event.Investor.String(),
				MessageSender:   event.Investor.String(),
				MessageAmount:   int(event.Amount.Int64()),
				MessagePosition: int(event.Position.Int64()),
				MessageNonce:    int(event.Nonce.Int64()),
				CreatedAt:       time.Now().UTC(),
				UpdatedAt:       time.Now().UTC(),
				ProcessStatus:   int(types.UnProcessed),
			}
			investEvents = append(investEvents, rawEvent)

		case backendabi.WithdrawEventSig:
			event := backendabi.PivotTopicWithdraw{}
			if err := utils.UnpackLog(backendabi.IPivotTopicABI, &event, "Withdraw", vlog); err != nil {
				log.Error("Failed to unpack Withdraw event", "err", err)
				return nil, nil, nil, nil, err
			}

			rawEvent := &orm.RawEvent{
				EventType:         int(types.Withdraw),
				ChainID:           types.Sepolia,
				ContractAddress:   types.ContractAddress,
				TxHash:            vlog.TxHash.String(),
				Timestamp:         uint64(time.Now().Unix()),
				BlockNumber:       vlog.BlockNumber,
				MessageWithdrawTo: event.To.String(),
				MessageAmount:     int(event.Amount.Int64()),
				MessageNonce:      int(event.Nonce.Int64()),
				CreatedAt:         time.Now().UTC(),
				UpdatedAt:         time.Now().UTC(),
				ProcessStatus:     int(types.UnProcessed),
			}
			withdrawEvents = append(withdrawEvents, rawEvent)

		case backendabi.WithdrawCommissionEventSig:
			event := backendabi.PivotTopicWithdrawCommission{}
			if err := utils.UnpackLog(backendabi.IPivotTopicABI, &event, "WithdrawCommission", vlog); err != nil {
				log.Error("Failed to unpack WithdrawCommission event", "err", err)
				return nil, nil, nil, nil, err
			}

			rawEvent := &orm.RawEvent{
				EventType:       int(types.WithdrawCommission),
				ChainID:         types.Sepolia,
				ContractAddress: types.ContractAddress,
				TxHash:          vlog.TxHash.String(),
				Timestamp:       uint64(time.Now().Unix()),
				BlockNumber:     vlog.BlockNumber,
				MessageSender:   event.Owner.String(),
				MessageAmount:   int(event.Amount.Int64()),
				MessageNonce:    int(event.Nonce.Int64()),
				CreatedAt:       time.Now().UTC(),
				UpdatedAt:       time.Now().UTC(),
				ProcessStatus:   int(types.UnProcessed),
			}
			withdrawCommissionEvents = append(withdrawCommissionEvents, rawEvent)

		}
	}
	logrus.Info("ParseSingleRawBridgeEventLogs ", " createTopicEvents ", len(createTopicEvents), " investEvents ", len(investEvents), " withdrawEvents ", len(withdrawEvents), " withdrawCommissionEvents ", len(withdrawCommissionEvents))
	return createTopicEvents, investEvents, withdrawEvents, withdrawCommissionEvents, nil
}

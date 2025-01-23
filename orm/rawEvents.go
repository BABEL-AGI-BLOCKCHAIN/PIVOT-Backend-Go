package orm

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// RawEvent represents a raw event.
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

func (RawEvent) TableName() string {
	return "raw_events"
}

// NewRawEvents creates a new instance of RawEvent.
func NewRawEvents(db *gorm.DB) *RawEvent {
	if err := db.AutoMigrate(&RawEvent{}); err != nil {
		panic(fmt.Sprintf("failed to auto migrate RawEvent, error: %v", err))
	}
	return &RawEvent{db: db}
}

// InsertRawEvents inserts a new RawEvent record into the database.
func (b *RawEvent) InsertRawEvents(ctx context.Context, tableName string, rawEvents []*RawEvent) error {
	fmt.Printf("Inserting %d raw events into table %s\n", len(rawEvents), tableName)
	if len(rawEvents) == 0 {
		return nil
	}
	db := b.db
	db = db.WithContext(ctx)
	db = db.Model(&RawEvent{})
	db = db.Table(tableName)
	// 'tx_status' column is not explicitly assigned during the update to prevent a later status from being overwritten back to "sent".

	for _, event := range rawEvents {
		if err := db.Create(event).Error; err != nil {
			if isDuplicateEntryError(err) {
				fmt.Printf("MessageNonce with hash %s already exists, skipping insert.\n", event.MessageNonce)
				continue
			}
			return fmt.Errorf("failed to insert message, error: %w", err)
		}
	}
	return nil
}
func isDuplicateEntryError(err error) bool {
	return strings.Contains(err.Error(), "Error 1062")
}

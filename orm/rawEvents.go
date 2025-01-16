package orm

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// RawEvents represents a raw event.
type RawEvents struct {
	db              *gorm.DB   `gorm:"column:-"`
	ID              uint64     `json:"id" gorm:"column:id;primary_key;autoIncrement"` // primary key in the database
	EventType       int        `json:"event_type" gorm:"column:event_type"`           //
	ChainID         int        `json:"chain_id" gorm:"column:chain_id"`               //
	ContractAddress string     `json:"contract_address" gorm:"column:contract_address"`
	TokenType       int        `json:"token_type" gorm:"column:token_type"`
	TxHash          string     `json:"tx_hash" gorm:"column:tx_hash"`
	GasPriced       string     `json:"gas_priced" gorm:"column:gas_priced"`
	BlockNumber     uint64     `json:"block_number" gorm:"column:block_number"`
	GasUsed         uint64     `json:"gas_used" gorm:"column:gas_used"`
	MsgValue        string     `json:"msg_value" gorm:"column:msg_value"`
	Timestamp       uint64     `json:"timestamp" gorm:"column:timestamp"`
	Sender          string     `json:"sender" gorm:"column:sender"` // sender address
	Receiver        string     `json:"receiver" gorm:"column:receiver"`
	MessageHash     string     `json:"message_hash" gorm:"column:message_hash;type:varchar(256);uniqueIndex"` // unique message hash
	CreatedAt       time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt       *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Remark          string     `json:"remark" gorm:"column:remark"`
	Status          int        `json:"status" gorm:"column:status"`
}

// NewRawEvents creates a new instance of RawEvents.
func NewRawEvents(db *gorm.DB) *RawEvents {

	return &RawEvents{db: db}
}

// InsertRawEvents inserts a new RawEvents record into the database.
func (b *RawEvents) InsertRawEvents(ctx context.Context, tableName string, rawEvents []*RawEvents) error {
	if len(rawEvents) == 0 {
		return nil
	}
	db := b.db
	db = db.WithContext(ctx)
	db = db.Model(&RawEvents{})
	db = db.Table(tableName)
	// 'tx_status' column is not explicitly assigned during the update to prevent a later status from being overwritten back to "sent".

	for _, event := range rawEvents {
		if err := db.Create(event).Error; err != nil {
			if isDuplicateEntryError(err) {
				fmt.Printf("Message with hash %s already exists, skipping insert.\n", event.MessageHash)
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

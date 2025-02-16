package logic

import (
	"context"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/orm"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
)

// EventUpdateLogic the logic of insert/update the database
type EventUpdateLogic struct {
	db           *gorm.DB
	rawEventsOrm *orm.RawEvent
}

// NewEventUpdateLogic creates a EventUpdateLogic instance
func NewEventUpdateLogic(db *gorm.DB) *EventUpdateLogic {
	b := &EventUpdateLogic{
		db:           db,
		rawEventsOrm: orm.NewRawEvents(db),
	}
	return b
}

// GetSyncHeight gets the sync height from db
func (b *EventUpdateLogic) GetSyncHeight(ctx context.Context) (uint64, error) {
	messageSyncedHeight, err := b.rawEventsOrm.GetMaxBlockNumber(ctx)
	if err != nil {
		log.Error("failed to get Max sync height", "err", err)
		return 0, err
	}

	return messageSyncedHeight, nil
}
func (b *EventUpdateLogic) InsertRawEvents(ctx context.Context, fetcherResult *FilterResult) error {
	if err := b.rawEventsOrm.InsertRawEvents(ctx, fetcherResult.createTopicEvents); err != nil {
		log.Error("failed to insert create topic events", "err", err)
		return err
	}
	if err := b.rawEventsOrm.InsertRawEvents(ctx, fetcherResult.investEvents); err != nil {
		log.Error("failed to insert invest events", "err", err)
		return err
	}
	if err := b.rawEventsOrm.InsertRawEvents(ctx, fetcherResult.withdrawEvents); err != nil {
		log.Error("failed to insert withdraw events", "err", err)
		return err
	}
	if err := b.rawEventsOrm.InsertRawEvents(ctx, fetcherResult.withdrawCommissionEvents); err != nil {
		log.Error("failed to insert withdraw commission events", "err", err)
		return err
	}
	return nil
}

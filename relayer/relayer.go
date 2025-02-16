package relayer

import (
	"context"
	"log"
	"time"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/orm"
	"gorm.io/gorm"
)

type Relayer struct {
	ctx              context.Context
	relayerSemaphore chan struct{}
	rawEventOrm      *orm.RawEvent
}

func NewRelayer(ctx context.Context, db *gorm.DB) (*Relayer, error) {
	relayer := &Relayer{
		ctx:              ctx,
		relayerSemaphore: make(chan struct{}, 1),
		rawEventOrm:      orm.NewRawEvents(db),
	}

	return relayer, nil
}
func (r *Relayer) Start() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			select {
			case r.relayerSemaphore <- struct{}{}:
				go func() {
					defer func() { <-r.relayerSemaphore }()
					r.pollingUnProcessedEvents()
				}()
			default:
				// skip this round if semaphore is full
			}
		case <-r.ctx.Done():
			return
		}
	}
}

func (r *Relayer) pollingUnProcessedEvents() {
	ctx := context.Background()
	unProcessedEvents, err := r.rawEventOrm.GetUnProcessedEvents(ctx, 500)
	if err != nil {
		log.Printf("Failed to query unconsumed messages: %v", err)
		return
	}
	for _, event := range unProcessedEvents {
		if event.EventType == int(orm.EventTypeCreateTopic) {
			//TODO
			// handle create topic event
			// [Table:Topic] insert
			// if(userId还不存在){
			// 	[Table:User] insert
			// }
		} else if event.EventType == int(orm.EventTypeInvest) {
			// if(userId还不存在){
			// 	[Table:User] insert
			// }
			// [Table:Investment] insert
			// [Table:Dividend] insert //本Message的用户的分红比例
			// [Table:Dividend] update //更新当前Message的TopicId中其它用户的分红比例

		} else if event.EventType == int(orm.EventTypeWithdraw) {
			//	[Table:Withdraw] insert

		} else if event.EventType == int(orm.EventTypeWithdrawCommission) {
			// handle withdraw commission event
		}
	}
}

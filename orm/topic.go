package orm

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Topic struct {
	db                *gorm.DB  `gorm:"column:-"`
	TopicID           int       `json:"topic_id" gorm:"column:topic_id;primary_key"` // primary key in the database
	TopicHash         string    `json:"topic_hash" gorm:"column:topic_hash"`
	Title             string    `json:"title" gorm:"column:title"`
	Resource          string    `json:"resource" gorm:"column:resource"`
	ResourceMimeType  string    `json:"resource" gorm:"column:resource_mime_type"`
	TotalInvestAmount float64   `json:"total_invest_amount" gorm:"column:total_invest_amount"`
	LatestPosition    float64   `json:"latest_position" gorm:"column:latest_position"`
	Content           string    `json:"content" gorm:"column:content"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at"`
	Promoter          string    `json:"creator" gorm:"column:promoter"`
	FixedInvestment   int       `json:"fixedInvestment" gorm:"column:fixed_Investment"`
}

type Investment struct {
	db                 *gorm.DB `gorm:"column:-"`
	ID                 int      `json:"id" gorm:"column:id;primary_key;autoIncrement"` // primary key in the database
	TopicID            int      `json:"topic_id" gorm:"column:topic_id"`
	UserID             string   `json:"user_id" gorm:"column:user_id"`
	Position           float64  `json:"position" gorm:"column:position"`
	InvestTokenAddress string   `json:"invest_token_address" gorm:"column:invest_token_address"`
	InvestTokenSymbol  string   `json:"invest_token_symbol" gorm:"column:invest_token_symbol"`
	InvestAmount       float64  `json:"invest_amount" gorm:"column:invest_amount"`
}

type Dividend struct {
	db                 *gorm.DB `gorm:"column:-"`
	ID                 int      `json:"id" gorm:"column:id;primary_key;autoIncrement"` // primary key in the database
	TopicID            int      `json:"topic_id" gorm:"column:topic_id"`
	UserID             string   `json:"user_id" gorm:"column:user_id"`
	DividendableAmount float64  `json:"dividendable_amount" gorm:"column:dividendable_amount"`
}

type Comment struct {
	db              *gorm.DB `gorm:"column:-"`
	CommentID       int      `json:"comment_id" gorm:"column:comment_id;primary_key;autoIncrement"` // primary key in the database
	Content         string   `json:"content" gorm:"column:content"`
	Creator         string   `json:"creator" gorm:"column:creator"` // commenter's userID
	TopicID         int      `json:"topic_id" gorm:"column:topic_id"`
	FatherCommentID int      `json:"father_comment_id" gorm:"column:father_comment_id"`
}

type Withdraw struct {
	db                *gorm.DB  `gorm:"column:-"`
	ID                int       `json:"id" gorm:"column:id;primary_key;autoIncrement"` // primary key in the database
	UserID            string    `json:"user_id" gorm:"column:user_id"`
	TopicID           int       `json:"topic_id" gorm:"column:topic_id"`
	WithdrawAmount    float64   `json:"withdraw_amount" gorm:"column:withdraw_amount"`
	WithdrawTimestamp time.Time `json:"withdraw_timestamp" gorm:"column:withdraw_timestamp"`
}

// NewTopic creates a new Topic instance
func NewTopic(db *gorm.DB) *Topic {
	if err := db.AutoMigrate(&Topic{}); err != nil {
		panic(fmt.Sprintf("failed to auto migrate Topic table, error: %v", err))
	}
	return &Topic{db: db}
}

// NewInvestment creates a new Investment instance
func NewInvestment(db *gorm.DB) *Investment {
	if err := db.AutoMigrate(&Investment{}); err != nil {
		panic(fmt.Sprintf("failed to auto migrate Investment table, error: %v", err))
	}
	return &Investment{db: db}
}

// NewDividend creates a new Dividend instance
func NewDividend(db *gorm.DB) *Dividend {
	if err := db.AutoMigrate(&Dividend{}); err != nil {
		panic(fmt.Sprintf("failed to auto migrate Dividend table, error: %v", err))
	}
	return &Dividend{db: db}
}

// NewComment creates a new Comment instance
func NewComment(db *gorm.DB) *Comment {
	if err := db.AutoMigrate(&Comment{}); err != nil {
		panic(fmt.Sprintf("failed to auto migrate Comment table, error: %v", err))
	}
	return &Comment{db: db}
}

// NewWithdraw creates a new Withdraw instance
func NewWithdraw(db *gorm.DB) *Withdraw {
	if err := db.AutoMigrate(&Withdraw{}); err != nil {
		panic(fmt.Sprintf("failed to auto migrate Withdraw table, error: %v", err))
	}
	return &Withdraw{db: db}
}

package api

import (
	"sync"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/controller/api/read"
	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/controller/api/write"
	"gorm.io/gorm"
)

var (
	// RPC Controllers
	GetUserInfoCtl      *read.GetUserInfoController
	GetUserInfoByCtl    *read.GetUserInfoByController
	TopicListCtl        *read.TopicListController
	GetTopicDetailByCtl *read.GetTopicDetailByController
	UserSignCtl         *write.UserSignController
	BindTwitterCtl      *write.BindTwitterController
	PostCommentCtl      *write.PostCommentController
	PostTopicCtl        *write.PostTopicController

	initControllerOnce sync.Once
)

// InitController inits Controller with database
func InitController(db *gorm.DB) {
	initControllerOnce.Do(func() {
		// Initialize RPC Controllers
		GetUserInfoCtl = read.NewGetUserInfoController(db)
		GetUserInfoByCtl = read.NewGetUserInfoByController(db)
		TopicListCtl = read.NewTopicListController(db)
		GetTopicDetailByCtl = read.NewGetTopicDetailByController(db)
		UserSignCtl = write.NewUserSignController(db)
		BindTwitterCtl = write.NewBindTwitterController(db)
		PostCommentCtl = write.NewPostCommentController(db)
		PostTopicCtl = write.NewPostTopicController(db)
	})
}

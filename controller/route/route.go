package route

import (
	"time"

	"github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go/controller/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Route routes the APIs
func Route(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("dev/")
	// bridgeApi := api.NewBridgeAPI(db)
	// r.GET("/test", bridgeApi.GetStatus)

	// Add RPC routes
	r.POST("/get_user_info", api.GetUserInfoCtl.HandleRequest)
	r.POST("/get_user_info_by", api.GetUserInfoByCtl.HandleRequest)
	r.POST("/topic_list", api.TopicListCtl.HandleRequest)
	r.POST("/get_topic_detail_by", api.GetTopicDetailByCtl.HandleRequest)
	r.POST("/user_sign", api.UserSignCtl.HandleRequest)
	r.POST("/bind_twitter", api.BindTwitterCtl.HandleRequest)
	r.POST("/post_comment", api.PostCommentCtl.HandleRequest)
	r.POST("/post_topic", api.PostTopicCtl.HandleRequest)
}

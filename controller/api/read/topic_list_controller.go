package read

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TopicListController struct {
	db *gorm.DB
}

func NewTopicListController(db *gorm.DB) *TopicListController {
	return &TopicListController{db: db}
}

func (ctl *TopicListController) HandleRequest(ctx *gin.Context) {
	// TODO: Implement the logic to handle the request
	// For initial testing, log a message and return a simple response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "TopicListController.HandleRequest called successfully",
	})
}

package read

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetTopicDetailByController struct {
	db *gorm.DB
}

func NewGetTopicDetailByController(db *gorm.DB) *GetTopicDetailByController {
	return &GetTopicDetailByController{db: db}
}

func (ctl *GetTopicDetailByController) HandleRequest(ctx *gin.Context) {
	// TODO: Implement the logic to handle the request
}

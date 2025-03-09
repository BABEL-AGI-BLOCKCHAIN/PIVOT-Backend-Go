package write

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostTopicController struct {
	db *gorm.DB
}

func NewPostTopicController(db *gorm.DB) *PostTopicController {
	return &PostTopicController{db: db}
}

func (ctl *PostTopicController) HandleRequest(ctx *gin.Context) {
	// TODO: Implement the logic to handle the request
}

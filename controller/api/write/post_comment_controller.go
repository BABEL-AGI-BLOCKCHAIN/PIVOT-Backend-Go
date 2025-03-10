package write

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostCommentController struct {
	db *gorm.DB
}

func NewPostCommentController(db *gorm.DB) *PostCommentController {
	return &PostCommentController{db: db}
}

func (ctl *PostCommentController) HandleRequest(ctx *gin.Context) {
	// TODO: Implement the logic to handle the request
}

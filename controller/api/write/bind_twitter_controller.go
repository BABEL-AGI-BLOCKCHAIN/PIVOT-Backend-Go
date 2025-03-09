package write

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BindTwitterController struct {
	db *gorm.DB
}

func NewBindTwitterController(db *gorm.DB) *BindTwitterController {
	return &BindTwitterController{db: db}
}

func (ctl *BindTwitterController) HandleRequest(ctx *gin.Context) {
	// TODO: Implement the logic to handle the request
}

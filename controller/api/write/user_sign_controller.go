package write

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserSignController struct {
	db *gorm.DB
}

func NewUserSignController(db *gorm.DB) *UserSignController {
	return &UserSignController{db: db}
}

func (ctl *UserSignController) HandleRequest(ctx *gin.Context) {
	// TODO: Implement the logic to handle the request
}

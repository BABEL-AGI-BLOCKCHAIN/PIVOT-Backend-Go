package read

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetUserInfoController struct {
	db *gorm.DB
}

func NewGetUserInfoController(db *gorm.DB) *GetUserInfoController {
	return &GetUserInfoController{db: db}
}

func (ctl *GetUserInfoController) HandleRequest(ctx *gin.Context) {
	// Implement the logic to handle the request
}

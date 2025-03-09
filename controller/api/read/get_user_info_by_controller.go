package read

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetUserInfoByController struct {
	db *gorm.DB
}

func NewGetUserInfoByController(db *gorm.DB) *GetUserInfoByController {
	return &GetUserInfoByController{db: db}
}

func (ctl *GetUserInfoByController) HandleRequest(ctx *gin.Context) {
	// TODO: Implement the logic to handle the request
}

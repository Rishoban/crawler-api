package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CrawlerService struct {
	DbConnection *gorm.DB `json:"dbConnection"`
}

func (v *CrawlerService) InitRouter(routerEngine *gin.Engine) {}

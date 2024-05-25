package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReponseProduct interface {
	HandlerCreateFood(db *gorm.DB) gin.HandlerFunc
}

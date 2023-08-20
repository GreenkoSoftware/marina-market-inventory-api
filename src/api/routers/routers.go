package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/GreenkoSoftware/marina-market-inventory-api/src/api/middlewares"
	"github.com/GreenkoSoftware/marina-market-inventory-api/src/api/routers/repository"
)

func Setup(db *gorm.DB, gin *gin.Engine) {

	// Public Routes
	publicRouter := gin.Group("")
	protectedRouter := gin.Group("")

	//Add jwt middleware
	protectedRouter.Use(middlewares.JwtTokenCheck)

	repository.NewUserRoute(db, protectedRouter)
	repository.NewAuthRoute(db, publicRouter)
}

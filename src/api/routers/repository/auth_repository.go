package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	auth_controller "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/auth"
)

func NewAuthRoute(db *gorm.DB, group *gin.RouterGroup) {
	user := group.Group("auth")

	user.POST("login", func(c *gin.Context) {
		function.CreateResponse(
			auth_controller.Auth(c, db),
		)
	})
}

package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	user_controller "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/user"
)

func NewUserRoute(db *gorm.DB, group *gin.RouterGroup) {
	user := group.Group("user")

	user.POST("/create", func(c *gin.Context) {
		function.CreateResponse(
			user_controller.Create(c, db),
		)
	})

	//TODO:UPDATE USER
	user.PUT("", func(c *gin.Context) {
		function.CreateResponse(
			user_controller.PutBy(c, db),
		)
	})

	//TODO:GET USERS
	user.GET("", func(c *gin.Context) {
		function.CreateResponse(
			user_controller.Get(c, db),
		)
	})

	//TODO:DELETE USER
	user.DELETE("", func(c *gin.Context) {
		function.CreateResponse(
			user_controller.Delete(c, db),
		)
	})

	user.PUT("/password/reset", func(c *gin.Context) {
		function.CreateResponse(
			user_controller.ResetPassword(c, db),
		)
	})
}

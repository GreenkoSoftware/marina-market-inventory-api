package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	payment_controller "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/payment"
)

func NewPaymentRepository(db *gorm.DB, group *gin.RouterGroup) {
	user := group.Group("payment")
	//TODO:GET USERS
	user.GET("", func(c *gin.Context) {
		function.CreateResponse(
			payment_controller.Get(c, db),
		)
	})

}

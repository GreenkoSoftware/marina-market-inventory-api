package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	voucher_controller "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/voucher"
)

func NewVoucherRepository(db *gorm.DB, group *gin.RouterGroup) {
	user := group.Group("voucher")
	//TODO:GET USERS
	user.GET("", func(c *gin.Context) {
		function.CreateResponse(
			voucher_controller.Get(c, db),
		)
	})

}

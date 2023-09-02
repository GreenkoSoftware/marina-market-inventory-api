package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	saleticket "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/sale-ticket"
)

func NewSaleRepository(db *gorm.DB, group *gin.RouterGroup) {
	user := group.Group("sale-ticket")
	user.POST("create", func(c *gin.Context) {
		function.CreateResponse(
			saleticket.CreateSale(c, db),
		)
	})
}

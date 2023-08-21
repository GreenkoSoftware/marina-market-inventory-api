package repository

import (
	function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	product_controller "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductRoute( db *gorm.DB, group *gin.RouterGroup){
	// trunk-ignore(git-diff-check/error)
	product:= group.Group("product")	
	product.POST("/create",func( c *gin.Context){
		function.CreateResponse( product_controller.Create(c,db))
	})
}
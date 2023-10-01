package repository

import (
	function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	product_controller "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/product"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductRoute(db *gorm.DB, group *gin.RouterGroup) {
	// trunk-ignore(git-diff-check/error)
	product := group.Group("product")
	/* Create product */
	product.POST("/create", func(c *gin.Context) {
		function.CreateResponse(product_controller.Create(c, db))
	})
	/*  Delete product */
	product.DELETE("", func(c *gin.Context) {
		function.CreateResponse(
			product_controller.Delete(c, db),
		)
	})
	/* Get products */
	product.GET("", func(c *gin.Context) {
		function.CreateResponse(
			product_controller.Get(c, db),
		)
	})
	/* Get Categories */
	product.GET("/categories", func(c *gin.Context) {
		function.CreateResponse(
			product_controller.GetCategories(c, db),
		)
	})
	/* Create Category*/
	product.POST("/create/category", func(c *gin.Context) {
		function.CreateResponse(
			product_controller.CreateCategory(c, db),
		)
	})
	/* Get Type Stocks */
	product.GET("/type-stock", func(c *gin.Context) {
		function.CreateResponse(
			product_controller.GetTypeStocks(c, db),
		)
	})
	/* Update product */
	product.PUT("", func(c *gin.Context) {
		function.CreateResponse(
			product_controller.PutBy(c, db),
		)
	})
	/* Update stock */
	product.PUT("/stock", func(c *gin.Context) {
		function.CreateResponse(product_controller.PutStockBy(c, db))
	})

	//offer
	product.POST("/create/offer", func(c *gin.Context) {
		function.CreateResponse(product_controller.CreateOffer(c, db))
	})
	/* Get products */
	product.GET("/offer", func(c *gin.Context) {
		function.CreateResponse(
			product_controller.GetOffer(c, db),
		)
	})
	product.GET("/offer/:id", func(c *gin.Context) {
		function.CreateResponse(
			product_controller.GetOfferByID(c, db),
		)
	})

	product.DELETE("/offer/:id", func(c *gin.Context) {
		function.CreateResponse(
			product_controller.DeleteOffer(c, db),
		)
	})
}

package saleticket

import (
	constants "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/constant"
	"github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/sale-ticket/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSale(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		request models.Sale
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		return c, nil, err
	}

	//Discount stock

	for _, product := range request.SalesReceipt {
		err := sql.DiscountStock(db, product.ProductID, product.Quantity)
		if err != nil {
			return c, nil, err
		}
	}

	if err = sql.CreateSalesReceipt(db, request); err != nil {
		return c, &err, nil
	}
	return c, &constants.InsertSuccess, err
}

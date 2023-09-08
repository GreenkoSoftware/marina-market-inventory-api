package sql

import (
	"context"
	"time"

	constants "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/constant"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"gorm.io/gorm"
)

func CreateSalesReceipt(db *gorm.DB, sale models.Sale) (err error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	sale.Date = time.Now()
	results := db.
		WithContext(ctx).
		Create(&sale)

	if results.Error != nil {
		return results.Error
	}

	return nil
}

func DiscountStock(db *gorm.DB, productID int, amountToSubtract int) (err error) {

	var product models.Product
	if err := db.Preload("ProductStocks").First(&product, productID).Error; err != nil {
		return constants.ErrorProductNotExist
	}

	product.ProductStocks.Stock -= amountToSubtract
	if err := db.Save(&product.ProductStocks).Error; err != nil {
		return constants.ErrorProductNotExist
	}
	return
}

package sql

import (
	"context"
	"time"

	"github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product models.Product) (err error){
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	results := db.WithContext(ctx).Omit("ProductStockID").Create(&product)
	if results.Error != nil {
		return results.Error
	}
	product.ProductStock.ProductID = product.ID
	results = db.WithContext(ctx).Debug().Create(&product.ProductStock)

	if results.Error != nil {
		return results.Error
	}
	return nil
}

func Delete(db *gorm.DB, product models.Product) (err error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Delete(&product)

	if results.Error != nil {
		return results.Error
	}

	return nil
}

func Get(db *gorm.DB) (products *[]models.Product, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Preload("ProductCategory").
		Preload("StockType").
		WithContext(ctx).
		Find(&products).Error

	if results != nil {
		return nil, results
	}

	return products, nil
}
func GetByParam(db *gorm.DB, fiel string, value string) (products *[]models.Product, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Scopes(sql.By(fiel, value)).
		Preload("ProductCategory").
		Preload("StockType").
		WithContext(ctx).
		Find(&products).Error

	if results != nil {
		return nil, results
	}

	return products, nil
}
func PutBy(db *gorm.DB, field string, value string, product *models.Product) (err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Preload("ProductCategory").
		Preload("StockType").
		WithContext(ctx).
		Model(&product).
		Update(field, value).Error

	if results != nil {
		return results
	}

	return nil
}
func CreateProductStock(db *gorm.DB, stock models.ProductStock) (err error){
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	results := db.WithContext(ctx).Create(&stock)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
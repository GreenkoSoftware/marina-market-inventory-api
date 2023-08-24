package sql

import (
	"context"
	"time"

	"github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/GreenkoSoftware/marina-market-inventory-api/src/utils"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product models.Product) (err error){
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	results := db.WithContext(ctx).Omit("ProductStockID").Create(&product)
	if results.Error != nil {
		return results.Error
	}
	product.ProductStocks.ProductID = product.ID
	results = db.WithContext(ctx).Debug().Create(&product.ProductStocks)

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

func GetProduct(db *gorm.DB) (P *[]models.Product, err error) {

	var Products *[]models.Product
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Preload("ProductCategories").
		Preload("StockTypes").
		Preload("ProductStocks").
		WithContext(ctx).
		Find(&Products).Error

	if results != nil {
		return nil, results
	}
	if !utils.HasData(*Products) {
		return nil, nil
	}
	return Products, nil
}
/* 
Debug().
		WithContext(ctx).
		Table("products").
		Preload("ProductCategory").
		Preload("StockType").
		Joins("inner join product_stocks on products.ID = product_stocks.product_id").
*/



func GetCategories(db *gorm.DB) (categories *[]models.ProductCategories, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Find(&categories).Error

	if results != nil {
		return nil, results
	}

	return categories, nil
}
func GetTypeStocks(db *gorm.DB) (typeStocks *[]models.StockTypes, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Find(&typeStocks).Error

	if results != nil {
		return nil, results
	}

	return typeStocks, nil
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

func CreateProductStock(db *gorm.DB, stock models.ProductStocks) (err error){
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	results := db.WithContext(ctx).Create(&stock)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
/* db, stockValue,stockMin, ProductID */
func PutStockBy(db *gorm.DB, stockValue,stockMin,ProductID int) (err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)

	defer cancel()
	results := db.
		WithContext(ctx).
		Table("product_stocks").
		Where("product_id = ?",ProductID).
		Update("stock_min",stockMin).
		Update("stock",stockValue).Error

	if results != nil {
		return results
	}

	return nil
}
/* db,productName,costPrice,netPrice,image,code,productCategoryId, stockTypeId */
func PutBy(db *gorm.DB, productName string,costPrice,netPrice float64,image,code string,productCategoryId, stockTypeId,ProductID int) (err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Preload("ProductCategory").
		Preload("StockType").
		WithContext(ctx).
		Table("products").
		Where("id = ?",ProductID).
		Update("name",productName).
		Update("cost_price",costPrice).
		Update("net_price",netPrice).
		Update("image",image).
		Update("code",code).
		Update("product_category_id",productCategoryId).
		Update("stock_type_id",stockTypeId).Error

	if results != nil {
		return results
	}

	return nil
}
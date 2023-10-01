package sql

import (
	"context"
	"time"

	"github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/GreenkoSoftware/marina-market-inventory-api/src/utils"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product models.Product) (err error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	results := db.WithContext(ctx).Create(&product)
	if results.Error != nil {
		return results.Error
	}
	results = db.WithContext(ctx).Debug().Create(&models.ProductStocks{
		ProductID: product.ID,
		Stock:     product.ProductStocks.Stock,
		StockMin:  product.ProductStocks.StockMin,
	})

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
		Debug().
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
func GetProductByID(db *gorm.DB, ID int) (P *models.Product, err error) {

	var Products *models.Product
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Preload("ProductStocks").
		Preload("ProductCategories").
		Preload("StockTypes").
		Where("ID = ?", ID).
		Find(&Products).Error

	if results != nil {
		return nil, results
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

func CreateProductStock(db *gorm.DB, stock models.ProductStocks) (err error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	results := db.WithContext(ctx).Create(&stock)
	if results.Error != nil {
		return results.Error
	}
	return nil
}

/* db, stockValue,stockMin, ProductID */
func PutStockBy(db *gorm.DB, stockValue, stockMin, ProductID int) (err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)

	defer cancel()
	results := db.
		WithContext(ctx).
		Table("product_stocks").
		Where("product_id = ?", ProductID).
		Update("stock_min", stockMin).
		Update("stock", stockValue).Error

	if results != nil {
		return results
	}

	return nil
}

/* db,productName,costPrice,netPrice,image,code,productCategoryId, stockTypeId */
func PutBy(db *gorm.DB, ProductID int, product *models.Product, productStocks *models.ProductStocks, stockTypeId, categoriesTypeId int) (err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Table("products").
		Preload("ProductStocks").
		Preload("ProductCategories").
		Preload("StockTypes").
		Where("id = ?", ProductID).
		Save(&product).Error

	if results != nil {
		return results
	}
	results = db.
		WithContext(ctx).
		Table("products").
		Where("id = ?", ProductID).
		Update("product_categories_id", categoriesTypeId).
		Update("stock_types_id", stockTypeId).Error

	if results != nil {
		return results
	}
	results = db.
		WithContext(ctx).
		Table("product_stocks").
		Where("product_id = ?", ProductID).
		Save(&productStocks).Error

	if results != nil {
		return results
	}
	return nil
}

func CreateProductOffer(db *gorm.DB, product models.ProductOffer) (err error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	results := db.WithContext(ctx).Create(&product)
	if results.Error != nil {
		return results.Error
	}
	return nil
}

func GetProductOffer(db *gorm.DB) (P *[]models.ProductOffer, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Preload("Product").
		WithContext(ctx).
		Find(&P).Error

	if results != nil {
		return nil, results
	}
	if !utils.HasData(*P) {
		return nil, nil
	}
	return
}
func GetProductOfferByID(db *gorm.DB, id int) (P *[]models.ProductOffer, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Preload("Product").
		WithContext(ctx).
		Where("product_id = ?", id).
		Find(&P).Error

	if results != nil {
		return nil, results
	}
	if !utils.HasData(*P) {
		return nil, nil
	}
	return
}

func DeleteOffer(db *gorm.DB, product models.ProductOffer) (err error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Debug().
		Delete(&product)

	if results.Error != nil {
		return results.Error
	}

	return nil
}

func GetCategoriesByName(db *gorm.DB, categoryName string) (categories *models.ProductCategories, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 50*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Where("name = ?", categoryName).
		Find(&categories).Error

	if results != nil {
		return nil, results
	}

	return categories, nil
}
func CreateCategoryProduct(db *gorm.DB, category models.ProductCategories) (err error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	results := db.WithContext(ctx).Create(&category)
	if results.Error != nil {
		return results.Error
	}
	return nil
}

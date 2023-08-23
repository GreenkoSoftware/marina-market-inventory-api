package models

import (
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	CostPrice float64 `json:"cost_price" binding:"required"`
	NetPrice float64 `json:"net_price" binding:"required"`
	Image string `json:"image,omitempty"`
	Code string `json:"code" binding:"required"`
	ProductCategoryID int `json:"product_category_id" binding:"required"`
	ProductCategory ProductCategory `gorm:"foreignKey:ProductCategoryID"`
	StockTypeID int  `json:"stock_type_id" binding:"required"`
	StockType StockType `gorm:"foreignKey:StockTypeID"`
	ProductStock ProductStock `gorm:"-" json:"product_stock"`
}
type ProductCategory struct {
	gorm.Model
	CategoryName string  `gorm:"unique;not null" json:"category_name"`
}

type StockType struct {
	gorm.Model
	TypeNameStock string `json:"type_name_stock"`
}

type ProductStock struct {
	gorm.Model
	Stock int `json:"stock" binding:"required"`
	StockMin int  `json:"stock_min" binding:"required"`
	ProductID uint `gorm:"foreignKey,unique" json:"product_id"`
}
func (product *Product) NormalizedProduct() {
	product.Name = strings.ToLower(product.Name)
	product.ProductCategory.CategoryName = strings.ToLower(product.ProductCategory.CategoryName)
	product.StockType.TypeNameStock = strings.ToLower(product.StockType.TypeNameStock)
}

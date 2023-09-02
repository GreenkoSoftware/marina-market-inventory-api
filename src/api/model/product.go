package models

import (
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name                string             `json:"name" binding:"required"`
	CostPrice           int                `json:"cost_price" binding:"required"`
	NetPrice            int                `json:"net_price" binding:"required"`
	SalePrice           int                `json:"sale_price" binding:"required"`
	Image               string             `json:"image,omitempty"`
	Code                string             `gorm:"unique;not null" json:"code" binding:"required"`
	ProductCategories   *ProductCategories `gorm:"foreignKey:ProductCategoriesID;references:ID"`
	StockTypes          *StockTypes        `gorm:"foreignKey:StockTypesID;references:ID"`
	ProductCategoriesID int                `json:"product_categories_id"`
	StockTypesID        int                `json:"stock_types_id"`
	ProductStocks       ProductStocks      `json:"product_stock"`
}

type ProductCategories struct {
	gorm.Model
	Name string `gorm:"unique;not null" json:"name"`
}

type StockTypes struct {
	gorm.Model
	Name string `json:"name"`
}

type ProductStocks struct {
	gorm.Model
	Stock     int  `json:"stock"`
	StockMin  int  `json:"stock_min"`
	ProductID uint `gorm:"foreignKey" json:"product_id"`
}

func (product *Product) NormalizedProduct() {
	product.Name = strings.ToLower(product.Name)
}

type ProductOffer struct {
	gorm.Model
	Quantity  int      `json:"quantity" binding:"required"`
	UnitPrice int      `json:"unit_price" binding:"required"`
	ProductID int      `json:"product_id"`
	Product   *Product `json:"product"`
}

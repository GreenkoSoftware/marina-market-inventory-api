package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name"`
	CostPrice float64 `json:"cost_price"`
	NetPrice float64 `json:"net_price"`
	Image string `json:"image"`
	Code string `json:"code"`
	ProductCategoryID string `json:"-"`
	ProductCategory ProductCategory `gorm:"foreignKey:ProductCategoryID"`
	StockTypeID string `json:"-"`
	StockType StockType `gorm:"foreignKey:StockTypeID"`
	ProductStockID string `json:"-"`
	ProductStock ProductStock `gorm:"foreignKey:StockTypeID"`
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
	Stock int `json:"-"`
	StockMin int `json:"-"`
}

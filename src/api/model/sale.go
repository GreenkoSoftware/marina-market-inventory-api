package models

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	Date          time.Time      `json:"date"`
	VoucherTypeID int            `json:"voucher_type_id"`
	PaymentTypeID int            `json:"payment_type_id"`
	SalesReceipt  []SalesReceipt `json:"sales_receipt"`
	VoucherType   VoucherType    `gorm:"foreignKey:VoucherTypeID"`
	PaymentType   PaymentType    `gorm:"foreignKey:PaymentTypeID"`
}

type VoucherType struct {
	gorm.Model
	Name string `json:"name"`
}

type PaymentType struct {
	gorm.Model
	Name string `json:"name"`
}

type SalesReceipt struct {
	gorm.Model
	ProductID  int     `json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductID"`
	SaleID     int     `json:"sale_id"`
	Sale       Sale    `gorm:"foreignKey:SaleID"`
	Quantity   float64 `json:"quantity"`
	TotalPrice int     `json:"total_price"`
}

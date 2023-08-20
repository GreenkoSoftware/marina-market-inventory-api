package models

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	Date time.Time `json:"date"`
	VouchertTypeID int `json:"-"`
	VoucherType VoucherType `gorm:"foreignKey:VoucherType"`
	PaymentTypeId int `json:"-"`
	PaymentType PaymentType `gorm:"foreignKey:PaymentType"`
}
type SalesReceipt struct {
	gorm.Model
  	TotalPrice float64 `json:"total_price"`
  	Quantity float64 `json:"quantity"`
	ProductId int  `json:"-"`
	Product Product  `json:"product"`
	SaleId int  `json:"-"`
	Sale Sale  `json:"sale"`
}
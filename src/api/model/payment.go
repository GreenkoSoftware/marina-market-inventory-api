package models

import "gorm.io/gorm"

type PaymentType struct {
	gorm.Model
	NamePaymentType string `gorm:"unique;not null" json:"name_payment_type"`
}
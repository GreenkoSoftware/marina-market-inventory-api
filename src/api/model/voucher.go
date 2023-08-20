package models

import "gorm.io/gorm"

type VoucherType struct {
	gorm.Model
	TypeNameVoucher string  `gorm:"unique;not null" json:"type_name_voucher"`
}
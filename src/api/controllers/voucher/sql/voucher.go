package sql

import (
	"context"
	"time"

	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"gorm.io/gorm"
)

func Get(db *gorm.DB) (voucher *[]models.VoucherType, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Find(&voucher).Error

	if results != nil {
		return nil, results
	}

	return voucher, nil
}

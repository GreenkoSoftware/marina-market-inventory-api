package sql

import (
	"context"
	"time"

	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, product models.Product) (err error){
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	results := db.WithContext(ctx).Create(&product)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
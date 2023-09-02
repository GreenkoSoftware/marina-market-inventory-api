package user

import (
	sql_event "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/payment/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		payment *[]models.PaymentType
	)

	if payment, err = sql_event.Get(db); err != nil {
		return c, &err, nil
	}
	return c, payment, err

}

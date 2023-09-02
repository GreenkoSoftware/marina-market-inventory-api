package user

import (
	sql_event "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/voucher/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		voucher *[]models.VoucherType
	)

	if voucher, err = sql_event.Get(db); err != nil {
		return c, &err, nil
	}
	return c, voucher, err

}

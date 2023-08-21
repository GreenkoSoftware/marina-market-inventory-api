package product

import (
	constants "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/constant"
	sql_event "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/product/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var ( request models.Product)
	if err = c.ShouldBindJSON(&request);err!=nil{
		return c,nil,err
	}
	request.NormalizedProduct()
	if err =  sql_event.CreateProduct(db,request);err!=nil{
		return c, constants.InsertSuccess, err
	}else {
		return c, &err, nil
	}
}
package product

import (
	"strconv"

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
func Delete(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		request models.Product
	)

	ProductIdstr := c.Query("id")
	ProductID, err := strconv.Atoi(ProductIdstr)
	if err != nil {
		return c, nil, err
	}

	request.ID = uint(ProductID)
	if err = sql_event.Delete(db, request); err != nil {
		return c, nil, err
	} else {
		return c, &constants.DeleteSuccess, nil
	}
}

func Get(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		products *[]models.Product
	)

	if ProductIDstr := c.Query("id"); ProductIDstr != "" {
		if products, err := sql_event.GetByParam(db, "id", ProductIDstr); err != nil {
			return c, &err, nil
		} else {
			return c, products, err
		}
	}
	if products, err = sql_event.Get(db); err != nil {
		return c, &err, nil
	}
	return c, products, err

}
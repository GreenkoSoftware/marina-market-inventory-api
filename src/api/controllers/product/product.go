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

func PutBy(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		product *models.Product
	)

	if ProductIDstr := c.Query("id"); ProductIDstr != "" {

		ProductID, err := strconv.Atoi(ProductIDstr)
		if err != nil {
			return c, nil, err
		}

		if ProductIDstr := c.Query("id"); ProductIDstr != "" {
			var field, value string
			/* Editar campos de producto */
			if ProductName :=c.Query("name"); ProductName != "" {
				field = "name"
				value = ProductName
			}
			if CostPrice :=c.Query("cost_price"); CostPrice != "" {
				field = "cost_price"
				value = CostPrice
			}
			if NetPrice :=c.Query("net_price"); NetPrice != "" {
				field = "net_price"
				value = NetPrice
			}
			if Image :=c.Query("image"); Image != "" {
				field = "image"
				value = Image
			}
			if Code :=c.Query("code"); Code != "" {
				field = "code"
				value = Code
			}
			if ProductCategoryID :=c.Query("product_category_id"); ProductCategoryID != "" {
				field = "product_category_id"
				value = ProductCategoryID
			}
			if StockTypeID :=c.Query("stock_type_id"); StockTypeID != "" {
				field = "stock_type_id"
				value = StockTypeID
			}
			product = &models.Product{}
			product.ID = uint(ProductID)
			if err := sql_event.PutBy(db, field, value, product); err != nil {
				return c, &err, nil
			} else {
				return c, &constants.UpdateSuccess, err
			}
		}
	}

	return c, &err, nil

}
/* Create New Stock product */
func CreateStock(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var ( request models.ProductStock)
	if err = c.ShouldBindJSON(&request);err!=nil{
		return c,nil,err
	}
	if err =  sql_event.CreateProductStock(db,request);err!=nil{
		return c, constants.InsertSuccess, err
	}else {
		return c, &err, nil
	}
}
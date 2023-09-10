package product

import (
	"fmt"
	"strconv"

	constants "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/constant"
	sql_event "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/product/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		request models.Product
	)
	if err = c.ShouldBindJSON(&request); err != nil {
		return c, nil, err
	}
	request.NormalizedProduct()

	if err = sql_event.CreateProduct(db, request); err == nil {
		return c, &constants.InsertSuccess, err
	} else {
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
	if products, err = sql_event.GetProduct(db); err != nil {
		return c, &err, nil
	}
	return c, products, err
}

/* Get categories */
func GetCategories(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		categories *[]models.ProductCategories
	)

	if categories, err = sql_event.GetCategories(db); err != nil {
		return c, &err, nil
	}
	return c, categories, err
}

/* Get Type Stock */
func GetTypeStocks(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		typeStocks *[]models.StockTypes
	)
	if typeStocks, err = sql_event.GetTypeStocks(db); err != nil {
		return c, &err, nil
	}
	return c, typeStocks, err
}

/* Create New Stock product */
func CreateStock(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		request models.ProductStocks
	)
	if err = c.ShouldBindJSON(&request); err != nil {
		return c, nil, err
	}
	if err = sql_event.CreateProductStock(db, request); err != nil {
		return c, constants.InsertSuccess, err
	} else {
		return c, &err, nil
	}
}

/* Update Stock product */
func PutStockBy(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {

	if ProductIDstr := c.Query("id"); ProductIDstr != "" {

		ProductID, err := strconv.Atoi(ProductIDstr)
		if err != nil {
			return c, nil, err
		}
		if ProductIDstr := c.Query("id"); ProductIDstr != "" {
			var stockValue, stockMin int
			/* Editar campos de producto */
			if Stock := c.Query("stock_max"); Stock != "" {
				stockValue, err = strconv.Atoi(Stock)
				if err != nil {
					return c, nil, err
				}
			}
			if StockMin := c.Query("stock_min"); StockMin != "" {
				stockMin, err = strconv.Atoi(StockMin)
				if err != nil {
					return c, nil, err
				}
			}
			if err := sql_event.PutStockBy(db, stockValue, stockMin, ProductID); err != nil {
				return c, &err, nil
			} else {
				return c, &constants.UpdateSuccess, err
			}
		}
	}

	return c, &err, nil

}

/* Update Product */
func PutBy(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		productStocks *models.ProductStocks
	)

	if ProductIDstr := c.Query("id"); ProductIDstr != "" {

		ProductID, err := strconv.Atoi(ProductIDstr)
		if err != nil {
			return c, nil, err
		}

		if ProductIDstr := c.Query("id"); ProductIDstr != "" {

			Product, err := sql_event.GetProductByID(db, ProductID)
			if err != nil {
				return c, nil, constants.ErrorProductNotExist
			}

			/* Editar campos de producto */
			if ProductName := c.Query("name"); ProductName != "" {
				Product.Name = ProductName
			}
			if CostPrice := c.Query("cost_price"); CostPrice != "" {
				Product.CostPrice, err = strconv.Atoi(CostPrice)
				if err != nil {
					return c, nil, err
				}
			}
			if NetPrice := c.Query("net_price"); NetPrice != "" {
				Product.NetPrice, err = strconv.Atoi(NetPrice)
				if err != nil {
					return c, nil, err
				}
			}
			if SalePrice := c.Query("sale_price"); SalePrice != "" {
				Product.SalePrice, err = strconv.Atoi(SalePrice)
				if err != nil {
					return c, nil, err
				}
			}
			if Image := c.Query("image"); Image != "" {
				Product.Image = Image
			}
			if Code := c.Query("code"); Code != "" {
				Product.Code = Code
			}
			if StockMin := c.Query("stock_min"); StockMin != "" {
				Product.ProductStocks.StockMin, err = strconv.ParseFloat(StockMin, 64)
				if err != nil {
					return c, nil, err
				}
			}
			if Stock := c.Query("stock"); Stock != "" {
				Product.ProductStocks.Stock, err = strconv.ParseFloat(Stock, 64)
				if err != nil {
					return c, nil, err
				}
			}
			if ProductCategoryID := c.Query("product_category_id"); ProductCategoryID != "" {
				Product.ProductCategoriesID, err = strconv.Atoi(ProductCategoryID)
				if err != nil {
					return c, nil, err
				}
			}
			if StockTypeID := c.Query("stock_type_id"); StockTypeID != "" {
				Product.StockTypesID, err = strconv.Atoi(StockTypeID)
				if err != nil {
					return c, nil, err
				}
			}
			productStocks = &Product.ProductStocks

			if err := sql_event.PutBy(db, ProductID, Product, productStocks, Product.StockTypesID, Product.ProductCategoriesID); err != nil {
				return c, &err, nil
			} else {
				return c, &constants.UpdateSuccess, err
			}
		}
	}

	return c, &err, nil

}

func CreateOffer(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		request models.ProductOffer
	)

	err = c.ShouldBindJSON(&request)
	if err != nil {
		return c, nil, err
	}

	_, err = sql_event.GetProductByID(db, request.ProductID)
	if err != nil {
		return c, nil, constants.ErrorProductNotExist
	}

	err = sql_event.CreateProductOffer(db, request)
	if err != nil {
		return c, nil, err
	}

	return c, &constants.InsertSuccess, nil
}

func GetOffer(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		products *[]models.ProductOffer
	)

	if products, err = sql_event.GetProductOffer(db); err != nil {
		return c, &err, nil
	}
	return c, products, nil
}

func GetOfferByID(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		products *[]models.ProductOffer
	)
	ProductIDstr := c.Params.ByName("id")
	if ProductIDstr == "" {
		return c, nil, constants.ErrorProductNotProductID
	}

	id, err := strconv.Atoi(ProductIDstr)
	if err != nil {
		fmt.Println(err)
		return
	}

	if products, err = sql_event.GetProductOfferByID(db, id); err != nil {
		return c, &err, nil
	}
	return c, products, err
}

func DeleteOffer(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		request models.ProductOffer
	)

	ProductIdstr := c.Params.ByName("id")
	ProductID, err := strconv.Atoi(ProductIdstr)
	if err != nil {
		return c, nil, err
	}

	request.ID = uint(ProductID)
	if err = sql_event.DeleteOffer(db, request); err != nil {
		return c, nil, err
	} else {
		return c, &constants.DeleteSuccess, nil
	}
}

//transform string into int golang?

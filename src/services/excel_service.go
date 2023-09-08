package services

import (
	"log"
	"strconv"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/product/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/sirupsen/logrus"
)

// ExcelService ExcelService
type ExcelService struct {
	logger            *logrus.Entry
	postgreSQLService *PostgreSQLService
}

// NewExcelService returns a service instance.
func NewExcelService() *ExcelService {
	return &ExcelService{}
}

// Health Health
func (service *ExcelService) Health() bool {
	return true
}

// InjectServices InjectServices
func (service *ExcelService) InjectServices(logger *logrus.Entry, services []Service) {
	service.logger = logger
	for _, otherService := range services {
		if PostgreSQLService, ok := otherService.(*PostgreSQLService); ok {
			service.postgreSQLService = PostgreSQLService
		}
	}
}

// Init Init this service
func (service *ExcelService) Init() error {
	service.logger.Info("[ExcelService] Initializing...")
	return nil
}

// Execute Execute this service
func (service *ExcelService) Execute(waitGroup *sync.WaitGroup) error {
	service.logger.Info("[ExcelService] Executing...")
	// Abre el archivo XLS
	archivo, err := excelize.OpenFile("dara.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// Nombre de la hoja que contiene los datos
	hoja := "Hoja 1"

	// Lee las filas del archivo XLS
	filas := archivo.GetRows(hoja)
	if err != nil {
		log.Fatal(err)
	}

	//var productos []models.Product

	for i, fila := range filas {
		if i != 0 {
			producto := models.Product{}
			for j, valor := range fila {
				switch j {
				case 0:
					producto.Code = valor
				case 1:
					if valor == "" {
						log.Print("aaa")
					}
					producto.Name = valor
				case 2:
					SalePrice, err := strconv.ParseFloat(valor, 64)
					producto.SalePrice = int(SalePrice)
					if err != nil {
						log.Fatalf("Error al convertir PrecioVenta en la fila %d: %v", i+1, err)
					}
				case 3:
					productoCategory, _ := sql.GetCategoriesByName(service.postgreSQLService.db, valor)
					producto.ProductCategoriesID = int(productoCategory.ID)
				case 4:
					CostPrice, err := strconv.ParseFloat(valor, 64)
					producto.CostPrice = int(CostPrice)
					if err != nil {
						producto.CostPrice = 0
					}
				}
			}
			productStock := models.ProductStocks{
				ProductID: producto.ID,
				StockMin:  0,
				Stock:     0,
			}
			producto.ProductStocks = productStock
			producto.StockTypesID = 2
			sql.CreateProduct(service.postgreSQLService.db, producto)
		}

	}

	return nil
}

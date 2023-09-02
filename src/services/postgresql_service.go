package services

import (
	"fmt"
	"log"
	"sync"

	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// postgreSQLService postgreSQLService
type PostgreSQLService struct {
	connectionStringCop string
	logger              *logrus.Entry
	db                  *gorm.DB
}

// NewpostgreSQLService returns a service instance.
func NewPostgreSQLService(host, port, name, user, pass string) *PostgreSQLService {

	copConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, pass, name)
	return &PostgreSQLService{
		connectionStringCop: copConnectionString,
	}
}

// Health Health
func (service *PostgreSQLService) Health() bool {
	return true
}

// InjectServices InjectServices
func (service *PostgreSQLService) InjectServices(logger *logrus.Entry, otherServices []Service) {
	service.logger = logger
}

// Init Init this service
func (service *PostgreSQLService) Init() error {
	service.logger.Info("[postgreSQLService] Initializing...")
	service.logger.Info("[postgreSQLService] Using connection string: " + service.connectionStringCop)

	database, err := gorm.Open(postgres.Open(service.connectionStringCop), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		log.Fatal("[postgreSQLService] Could not load environment variables", err)
	}
	service.db = database
	//Si estas estructura no existen en la bd, crealas como tabla.
	// trunk-ignore(golangci-lint/errcheck)
	service.db.AutoMigrate(
		&models.Product{},
		&models.ProductStocks{},
		&models.StockTypes{},
		&models.ProductCategories{},
		&models.ProductOffer{},
		&models.VoucherType{},
		&models.PaymentType{},
		&models.SalesReceipt{},
		&models.Sale{},
	)
	return nil
}

// Execute Execute this service
func (service *PostgreSQLService) Execute(waitGroup *sync.WaitGroup) error {
	service.logger.Info("[postgreSQLService] Executing...")

	waitGroup.Wait()
	return nil
}

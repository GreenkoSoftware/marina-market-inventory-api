package main

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/GreenkoSoftware/marina-market-inventory-api/src/services"
	"github.com/GreenkoSoftware/marina-market-inventory-api/src/utils"
)

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&utils.LogFormat{})

	logger := logrus.WithFields(nil)
	logger.Info("Initializing marina market inventory api...")

	allServices := []services.Service{

		services.NewPostgreSQLService(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASS"),
		),
		services.NewAPIService(
			os.Getenv("PORT"),
		),
	}
	for _, service := range allServices {
		service.InjectServices(logger, allServices)
	}

	for _, service := range allServices {
		// trunk-ignore(golangci-lint/errcheck)
		service.Init()
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(allServices))

	for _, service := range allServices {
		// trunk-ignore(golangci-lint/errcheck)
		go service.Execute(&waitGroup)
	}
	waitGroup.Wait()
}

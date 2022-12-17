package main

import (
	"fmt"

	"github.com/audi-skripsi/event_bus_migrator/internal/config"
	"github.com/audi-skripsi/event_bus_migrator/internal/repository"
	"github.com/audi-skripsi/event_bus_migrator/internal/service"
	"github.com/audi-skripsi/event_bus_migrator/pkg/util/logutil"
	"github.com/segmentio/kafka-go"
)

func main() {
	config.Init()
	config := config.Get()

	logger := logutil.NewLogger(logutil.NewLoggerParams{
		PrettyPrint: true,
		ServiceName: config.AppName,
	})

	logger.Infof("app initialized with the config of: %+v", config)

	kafkaAdmin, err := kafka.Dial("tcp", config.KafkaConfig.Address)
	if err != nil {
		logger.Fatalf("error getting kafka admin client: %+v", err)
		return
	}

	ctrl, err := kafkaAdmin.Controller()
	if err != nil {
		logger.Fatalf("error getting kafka controller: %+v", err)
		return
	}

	err = kafkaAdmin.Close()
	if err != nil {
		logger.Fatalf("error closing kafka admin client: %+v", err)
		return
	}

	kafkaAdmin, err = kafka.Dial("tcp", fmt.Sprintf("%s:%d", ctrl.Host, ctrl.Port))
	if err != nil {
		logger.Fatalf("error getting kafka admin client: %+v", err)
		return
	}

	repository := repository.NewRepository(repository.NewRepositoryParams{
		Logger:     logger,
		Config:     config,
		KafkaAdmin: kafkaAdmin,
	})

	service := service.NewService(service.NewServiceParams{
		Logger:     logger,
		Repository: repository,
		Config:     config,
	})

	err = service.MigrateTopics()
	if err != nil {
		logger.Fatalf("error migrating topics: %+v", err)
	}
	logger.Infof("topics migrated successfully")
}

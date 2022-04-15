package service

import (
	"github.com/audi-skripsi/event_bus_migrator/internal/config"
	"github.com/audi-skripsi/event_bus_migrator/internal/repository"
	"github.com/sirupsen/logrus"
)

type Service interface {
	MigrateTopics() (err error)
}

type service struct {
	logger     *logrus.Entry
	repository repository.Repository
	config     *serviceConfig
}

type serviceConfig struct {
	KafkaConfig *config.KafkaConfig
}

type NewServiceParams struct {
	Logger     *logrus.Entry
	Repository repository.Repository
	Config     *config.Config
}

func NewService(params NewServiceParams) Service {
	return &service{
		logger:     params.Logger,
		repository: params.Repository,
		config: &serviceConfig{
			KafkaConfig: &params.Config.KafkaConfig,
		},
	}
}

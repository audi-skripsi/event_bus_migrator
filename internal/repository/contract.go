package repository

import (
	"github.com/audi-skripsi/event_bus_migrator/internal/config"
	"github.com/audi-skripsi/event_bus_migrator/internal/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	MigrateTopics(topics []model.Topic) (err error)
}

type repository struct {
	logger     *logrus.Entry
	kafkaAdmin *kafka.AdminClient
	config     *repositoryConfig
}

type repositoryConfig struct {
	kafkaConfig config.KafkaConfig
}

type NewRepositoryParams struct {
	Logger     *logrus.Entry
	KafkaAdmin *kafka.AdminClient
	Config     *config.Config
}

func NewRepository(params NewRepositoryParams) Repository {
	return &repository{
		logger:     params.Logger,
		kafkaAdmin: params.KafkaAdmin,
		config: &repositoryConfig{
			kafkaConfig: params.Config.KafkaConfig,
		},
	}
}

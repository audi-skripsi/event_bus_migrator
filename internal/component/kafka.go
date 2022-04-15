package component

import (
	"github.com/audi-skripsi/event_bus_migrator/internal/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaAdminClient(config config.KafkaConfig) (client *kafka.AdminClient, err error) {
	client, err = kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": config.Address,
	})
	return
}

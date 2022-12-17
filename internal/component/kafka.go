package component

import (
	"github.com/audi-skripsi/event_bus_migrator/internal/config"
	"github.com/segmentio/kafka-go"
)

func NewKafkaAdminClient(config config.KafkaConfig) (client *kafka.Conn, err error) {
	client, err = kafka.Dial("tcp", config.Address)
	return
}

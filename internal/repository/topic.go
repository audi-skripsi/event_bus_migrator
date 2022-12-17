package repository

import (
	"github.com/audi-skripsi/event_bus_migrator/internal/model"
	"github.com/segmentio/kafka-go"
)

func (r *repository) MigrateTopics(topics []model.Topic) (err error) {
	var topicSpecs []kafka.TopicConfig
	var topicNames []string

	r.logger.Infof("migrating topics")

	for _, v := range topics {
		topicSpecs = append(topicSpecs, kafka.TopicConfig{
			Topic:             v.Name,
			NumPartitions:     v.Partition,
			ReplicationFactor: v.ReplicationFactor,
		})
		topicNames = append(topicNames, v.Name)
	}

	err = r.kafkaAdmin.DeleteTopics(topicNames...)
	if err != nil {
		r.logger.Errorf("error deleting topics: %+v", err)
	}

	err = r.kafkaAdmin.CreateTopics(
		topicSpecs...,
	)
	if err != nil {
		r.logger.Infof("error creating topic: %+v", err)
		return
	}
	r.logger.Infof("topics created")
	return
}

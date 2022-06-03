package repository

import (
	"context"

	"github.com/audi-skripsi/event_bus_migrator/internal/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (r *repository) MigrateTopics(topics []model.Topic) (err error) {
	var topicSpecs []kafka.TopicSpecification
	var topicNames []string

	for _, v := range topics {
		topicSpecs = append(topicSpecs, kafka.TopicSpecification{
			Topic:             v.Name,
			NumPartitions:     v.Partition,
			ReplicationFactor: v.ReplicationFactor,
		})
		topicNames = append(topicNames, v.Name)
	}

	res, err := r.kafkaAdmin.DeleteTopics(context.Background(), topicNames)
	if err != nil {
		r.logger.Errorf("error deleting topics: %+v", err)
		return
	}
	r.logger.Infof("delete topics response: %+v", res)

	res, err = r.kafkaAdmin.CreateTopics(context.Background(),
		topicSpecs,
	)
	if err != nil {
		r.logger.Infof("error creating topic: %+v", err)
		return
	}
	r.logger.Infof("topics created: %+v", res)
	return
}

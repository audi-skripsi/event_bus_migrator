package service

import (
	"encoding/json"
	"os"

	"github.com/audi-skripsi/event_bus_migrator/internal/dto"
	"github.com/audi-skripsi/event_bus_migrator/internal/model"
)

func (s *service) MigrateTopics() (err error) {
	topicMigrationPath := "./migration/topics.json"
	b, err := os.ReadFile(topicMigrationPath)
	if err != nil {
		s.logger.Errorf("error reading file: %+v", err)
		return
	}
	var topicMigration dto.TopicMigration
	err = json.Unmarshal(b, &topicMigration)
	if err != nil {
		s.logger.Errorf("error unmarshalling migration: %+v", err)
		return
	}
	var topicModels []model.Topic
	for _, v := range topicMigration.Topics {
		topicModels = append(topicModels, model.Topic{
			Name:              v.Name,
			Partition:         v.Partition,
			ReplicationFactor: v.ReplicationFactor,
		})
	}
	err = s.repository.MigrateTopics(topicModels)
	if err != nil {
		s.logger.Errorf("error migrating topics: %+v", err)
		return
	}
	return
}

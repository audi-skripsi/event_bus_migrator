package dto

type Topic struct {
	Name              string `json:"name"`
	Partition         int    `json:"partition"`
	ReplicationFactor int    `json:"replication_factor"`
}

type TopicMigration struct {
	Topics []Topic `json:"data"`
}

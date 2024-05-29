package hazelcast

type Config struct {
	Paths       []string `json:"paths"`
	ClusterName string   `json:"cluster_name"`
	MapName     string   `json:"map_name"`
	QueueName   string   `json:"queue_name"`
}

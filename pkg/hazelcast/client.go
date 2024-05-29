package hazelcast

import (
	"context"
	"github.com/hazelcast/hazelcast-go-client"
)

func NewHazelcast(cfg Config) (*hazelcast.Client, error) {
	config := hazelcast.Config{}

	config.Cluster.Network.SetAddresses(cfg.Paths...)
	config.Cluster.Name = cfg.ClusterName

	return hazelcast.StartNewClientWithConfig(context.Background(), config)
}

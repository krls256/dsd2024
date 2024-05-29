package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/hazelcast/hazelcast-go-client"
	"log"
)

func main() {
	config := hazelcast.Config{}

	config.Cluster.Network.SetAddresses("0.0.0.0:5701", "0.0.0.0:5702", "0.0.0.0:5703")
	config.Cluster.Name = "hzc"

	client, err := hazelcast.StartNewClientWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	nums, err := client.GetMap(context.Background(), "numbersv2")
	if err != nil {
		log.Fatal(err)
	}

	if err := nums.Clear(context.Background()); err != nil {
		log.Fatal(err)
	}

	for i := range 1000 {
		if err = nums.Set(context.Background(), i, uuid.New()); err != nil {
			log.Fatal(err)
		}
	}
}

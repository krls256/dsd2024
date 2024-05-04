package main

import (
	"context"
	"fmt"
	"github.com/hazelcast/hazelcast-go-client"
	"log"
	"sync"
)

var (
	config   = hazelcast.Config{}
	MaxValue = 100
	MaxSize  = 10
)

func init() {
	config.Cluster.Network.SetAddresses("0.0.0.0:5701", "0.0.0.0:5702", "0.0.0.0:5703")
	config.Cluster.Name = "hzc"
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go runWriter(wg)
	go runReader(wg, "reader 1:")
	go runReader(wg, "reader 2:")

	wg.Wait()
}

func runWriter(wg *sync.WaitGroup) {
	defer wg.Done()

	nums := getQueue("numbers_queue")

	nums.Clear(context.Background())

	for i := range MaxValue {
		if err := nums.Put(context.Background(), i); err != nil {
			log.Fatal(err)
		}
	}
}

func runReader(wg *sync.WaitGroup, readerName string) {
	defer wg.Done()

	nums := getQueue("numbers_queue")

	for range MaxValue {
		id, err := nums.Take(context.Background())
		fmt.Println(readerName, id)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func getQueue(name string) *hazelcast.Queue {
	client, err := hazelcast.StartNewClientWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	nums, err := client.GetQueue(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := nums.DrainWithMaxSize(context.Background(), MaxSize); err != nil {
		log.Fatal(err)
	}

	return nums
}

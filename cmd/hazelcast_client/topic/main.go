package main

import (
	"context"
	"fmt"
	"github.com/hazelcast/hazelcast-go-client"
	"log"
	"sync"
	"time"
)

var (
	config   = hazelcast.Config{}
	MaxValue = 100
)

func init() {
	config.Cluster.Network.SetAddresses("0.0.0.0:5701", "0.0.0.0:5702", "0.0.0.0:5703")
	config.Cluster.Name = "hzc"
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go runWriter(wg)
	go runReader(wg, "reader 1:", false)
	go runReader(wg, "reader 2:", true)

	wg.Wait()
}

func runWriter(wg *sync.WaitGroup) {
	defer wg.Done()

	nums := getTopic("numbers_topic")

	for i := range MaxValue {
		if err := nums.Publish(context.Background(), i); err != nil {
			log.Fatal(err)
		}
	}
}

func runReader(wg *sync.WaitGroup, readerName string, sleepInMiddle bool) {

	nums := getTopic("numbers_topic")

	i := 0

	id, err := nums.AddMessageListener(context.Background(), func(event *hazelcast.MessagePublished) {
		fmt.Println(readerName, event.Value)

		i++

		if sleepInMiddle && i == MaxValue/2 {
			time.Sleep(time.Second)
		}

		if i == MaxValue-1 {
			wg.Done()
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}

func getTopic(name string) *hazelcast.Topic {
	client, err := hazelcast.StartNewClientWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	nums, err := client.GetTopic(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}

	return nums
}

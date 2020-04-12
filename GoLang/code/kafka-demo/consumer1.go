package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.1.202:9092"}, nil)

	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions("testGo")

	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("testGo", int32(partition), sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
		wg.Wait()
		_ = consumer.Close()
		pc.AsyncClose()
	}
}
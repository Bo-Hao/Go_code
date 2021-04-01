package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func producer_test() {
	fmt.Printf("Producer\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer([]string{"0.0.0.0:50429"}, config)
	if err != nil {
		fmt.Printf("Create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic: "kafka_go_test",
		Key:   sarama.StringEncoder("go_test"),
	}

	value := "this is message"
	for {
		fmt.Scanln(&value)
		msg.Value = sarama.ByteEncoder(value)
		fmt.Printf("input [%s]\n", value)

		// send to chain
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}
}

func consumer_test() {
	fmt.Printf("Consumer")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"0.0.0.0:50429"}, config)
	if err != nil {
		fmt.Printf("Create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition("kafka_go_test", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}

}
func delete_topic() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	clusterAdmin, _ := sarama.NewClusterAdmin([]string{"0.0.0.0:50429"}, config)
	clusterAdmin.DeleteTopic("kafka_go_test1")
}

func metadata_test() {
	fmt.Printf("metadata test\n")

	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2

	client, err := sarama.NewClient([]string{"0.0.0.0:64881"}, config)
	if err != nil {
		fmt.Printf("metadata_test try create client err :%s\n", err.Error())
		return
	}

	defer client.Close()

	// get topic set
	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}

	fmt.Printf("topics(%d):\n", len(topics))

	for _, topic := range topics {
		fmt.Println(topic)
	}

	// get broker set
	brokers := client.Brokers()
	fmt.Printf("broker set(%d):\n", len(brokers))
	for _, broker := range brokers {
		fmt.Printf("%s\n", broker.Addr())
	}
}

func main() {
	//producer_test()
	//consumer_test()
	//metadata_test()
}

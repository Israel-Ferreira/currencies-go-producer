package producers

import (
	"encoding/json"

	"github.com/Israel-Ferreira/currencies-go-producer/domain"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type CurrenciesProducer struct {
	Producer *kafka.Producer
}

func (cp CurrenciesProducer) SendKafkaTicker(topic string, ticker domain.Ticker) error {

	msgValue, err := json.Marshal(ticker)

	if err != nil {
		return err
	}

	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msgValue,
	}

	// fmt.Println(config.Producer)

	if err = cp.Producer.Produce(&msg, nil); err != nil {
		return err
	}

	return nil

}

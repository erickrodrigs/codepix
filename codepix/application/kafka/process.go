package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	appmodel "github.com/erickrodrigs/codepix/codepix-go/application/appmodel"
	"github.com/erickrodrigs/codepix/codepix-go/application/factory"
	"github.com/erickrodrigs/codepix/codepix-go/application/usecase"
	"github.com/erickrodrigs/codepix/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

// MyKafkaProcessor ...
type MyKafkaProcessor struct {
	Database     *gorm.DB
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

// NewKafkaProcessor ...
func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *MyKafkaProcessor {
	return &MyKafkaProcessor{
		Database:     database,
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}

// Consume ...
func (kp *MyKafkaProcessor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "consumergroup",
		"auto.offset.reset": "earliest",
	}

	consumer, err := ckafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics := []string{"teste"}
	consumer.SubscribeTopics(topics, nil)

	fmt.Println("kafka consumer has been started")

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			fmt.Println(string(msg.Value))
			kp.processMessage(msg)
		}
	}
}

func (kp *MyKafkaProcessor) processMessage(msg *ckafka.Message) {
	transactionsTopic := "transactions"
	transactionConfirmationTopic := "transaction_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionsTopic:
		kp.processTransaction(msg)
	case transactionConfirmationTopic:
		kp.proccessTransactionConfirmation(msg)
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}
}

func (kp *MyKafkaProcessor) processTransaction(msg *ckafka.Message) error {
	transaction := appmodel.NewTransaction()
	err := transaction.ParseJSON(msg.Value)

	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(kp.Database)

	createdTransaction, err := transactionUseCase.Register(
		transaction.AccountID,
		transaction.Amount,
		transaction.PixKeyTo,
		transaction.PixKeyKindTo,
		transaction.Description,
	)

	if err != nil {
		fmt.Println("error registering transaction", err)
		return err
	}

	topic := "bank" + createdTransaction.PixKeyTo.Account.Bank.Code
	transaction.ID = createdTransaction.ID
	transaction.Status = model.TransactionPending
	transactionJSON, err := transaction.ToJSON()

	if err != nil {
		return err
	}

	err = Publish(string(transactionJSON), topic, kp.Producer, kp.DeliveryChan)

	if err != nil {
		return err
	}

	return nil
}

func (kp *MyKafkaProcessor) proccessTransactionConfirmation(msg *ckafka.Message) error {
	transaction := appmodel.NewTransaction()
	err := transaction.ParseJSON(msg.Value)

	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(kp.Database)

	if transaction.Status == model.TransactionConfirmed {
		err = kp.confirmTransaction(transaction, transactionUseCase)

		if err != nil {
			return err
		}
	} else if transaction.Status == model.TransactionCompleted {
		_, err := transactionUseCase.Complete(transaction.ID)

		if err != nil {
			return err
		}
	}

	return nil
}

func (kp *MyKafkaProcessor) confirmTransaction(transaction *appmodel.Transaction, transactionUseCase usecase.TransactionUseCase) error {
	confirmedTransaction, err := transactionUseCase.Confirm(transaction.ID)

	if err != nil {
		return err
	}

	topic := "bank" + confirmedTransaction.AccountFrom.Bank.Code
	transactionJSON, err := transaction.ToJSON()

	if err != nil {
		return err
	}

	err = Publish(string(transactionJSON), topic, kp.Producer, kp.DeliveryChan)

	if err != nil {
		return err
	}

	return nil
}

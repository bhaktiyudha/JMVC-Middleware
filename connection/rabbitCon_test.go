package connection

import (
	"errors"
	"strings"
	"testing"

	"github.com/NeowayLabs/wabbit"
)

type TestRabbitChannel struct {
	ExpectedError         error
	ExpectedErrorLocation string
	TestQueue             *TestRabbitQueue
}

type TestRabbitQueue struct {
	ExpectedResults map[string]interface{}
}

func (tc *TestRabbitChannel) Ack(tag uint64, multiple bool) error {
	if tc.ExpectedErrorLocation != "Ack" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) Nack(tag uint64, multiple bool, requeue bool) error {
	if tc.ExpectedErrorLocation != "Nack" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) Reject(tag uint64, requeue bool) error {
	if tc.ExpectedErrorLocation != "Reject" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) Confirm(noWait bool) error {
	if tc.ExpectedErrorLocation != "Confirm" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) NotifyPublish(confirm chan wabbit.Confirmation) chan wabbit.Confirmation {
	return nil
}

func (tc *TestRabbitChannel) Cancel(consumer string, noWait bool) error {
	if tc.ExpectedErrorLocation != "Cancel" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) ExchangeDeclare(name, kind string, opt wabbit.Option) error {
	if tc.ExpectedErrorLocation != "ExchangeDeclare" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) ExchangeDeclarePassive(name, kind string, opt wabbit.Option) error {
	if tc.ExpectedErrorLocation != "ExchangeDeclarePassive" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) QueueDeclare(name string, args wabbit.Option) (wabbit.Queue, error) {
	if tc.ExpectedErrorLocation != "QueueDeclare" {
		return tc.TestQueue, nil
	}
	return tc.TestQueue, tc.ExpectedError
}

func (tc *TestRabbitChannel) QueueDeclarePassive(name string, args wabbit.Option) (wabbit.Queue, error) {
	if tc.ExpectedErrorLocation != "QueueDeclarePassive" {
		return tc.TestQueue, nil
	}
	return tc.TestQueue, tc.ExpectedError
}

func (tc *TestRabbitChannel) QueueDelete(name string, args wabbit.Option) (int, error) {
	if tc.ExpectedErrorLocation != "QueueDelete" {
		return 0, nil
	}
	return 0, tc.ExpectedError
}

func (tc *TestRabbitChannel) QueueBind(name, key, exchange string, opt wabbit.Option) error {
	if tc.ExpectedErrorLocation != "QueueBind" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) QueueUnbind(name, route, exchange string, args wabbit.Option) error {
	if tc.ExpectedErrorLocation != "QueueUnbind" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) Consume(queue, consumer string, opt wabbit.Option) (<-chan wabbit.Delivery, error) {
	if tc.ExpectedErrorLocation != "Consume" {
		return nil, nil
	}
	return nil, tc.ExpectedError
}

func (tc *TestRabbitChannel) Qos(prefetchCount, prefetchSize int, global bool) error {
	if tc.ExpectedErrorLocation != "Qos" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) Close() error {
	if tc.ExpectedErrorLocation != "Close" {
		return nil
	}
	return tc.ExpectedError
}

func (tc *TestRabbitChannel) NotifyClose(chan wabbit.Error) chan wabbit.Error {
	return nil
}

func (tc *TestRabbitChannel) Publish(exc, route string, msg []byte, opt wabbit.Option) error {
	if tc.ExpectedErrorLocation != "Publish" {
		return nil
	}
	return tc.ExpectedError
}

func (tq *TestRabbitQueue) Name() string {
	value := ""

	if expectedValue, ok := tq.ExpectedResults["Name"]; ok {
		value = expectedValue.(string)
	}
	return value
}

func (tq *TestRabbitQueue) Messages() int {
	value := 0

	if expectedValue, ok := tq.ExpectedResults["Messages"]; ok {
		value = expectedValue.(int)
	}
	return value
}

func (tq *TestRabbitQueue) Consumers() int {
	value := 0

	if expectedValue, ok := tq.ExpectedResults["Consumers"]; ok {
		value = expectedValue.(int)
	}
	return value
}

func TestMakeQueue(t *testing.T) {
	var ch wabbit.Channel

	ch = &TestRabbitChannel{
		TestQueue: &TestRabbitQueue{},
	}

	_, err := MakeQueue(ch, "test")

	if err != nil {
		t.Fatal(err)
	}

	ch = &TestRabbitChannel{
		ExpectedErrorLocation: "QueueDeclare",
		ExpectedError:         errors.New("test_error"),
		TestQueue:             &TestRabbitQueue{},
	}

	_, err = MakeQueue(ch, "test_failed")

	if err == nil || (err != nil && !strings.Contains(err.Error(), "Error make query")) {
		t.Fatalf("Expected error %s but got %s", "Error make query", err)
	}
}

func TestMakeConsumer(t *testing.T) {
	var ch wabbit.Channel

	ch = &TestRabbitChannel{
		TestQueue: &TestRabbitQueue{
			ExpectedResults: map[string]interface{}{
				"Name": "test_consumer",
			},
		},
	}

	queue, err := MakeQueue(ch, "test_consumer")

	if err != nil {
		t.Fatal(err)
	}

	_, err = MakeConsumer(ch, "", queue)

	if err != nil {
		t.Fatal(err)
	}

	_, err = MakeConsumer(ch, "test_consumer_2")

	if err != nil {
		t.Fatal(err)
	}

	ch = &TestRabbitChannel{
		ExpectedErrorLocation: "QueueDeclare",
		ExpectedError:         errors.New("test error"),
		TestQueue: &TestRabbitQueue{
			ExpectedResults: map[string]interface{}{
				"Name": "test_consumer",
			},
		},
	}

	_, err = MakeConsumer(ch, "test_consumer_3")

	if err == nil || (err != nil && !strings.Contains(err.Error(), "Error make query")) {
		t.Fatalf("Expected error %s but got %s", "Error make query", err)
	}

	ch = &TestRabbitChannel{
		ExpectedErrorLocation: "Consume",
		ExpectedError:         errors.New("test error"),
		TestQueue: &TestRabbitQueue{
			ExpectedResults: map[string]interface{}{
				"Name": "test_consumer",
			},
		},
	}

	_, err = MakeConsumer(ch, "", queue)

	if err == nil || (err != nil && !strings.Contains(err.Error(), "Error build consumer for queue")) {
		t.Fatalf("Expected error %s but got %s", "Error build consumer for queue", err)
	}
}

package connection

import (
	"errors"
	"fmt"

	"github.com/NeowayLabs/wabbit"
)

func MakeQueue(ch wabbit.Channel, queueName string) (wabbit.Queue, error) {
	queue, err := ch.QueueDeclare(
		queueName,
		nil,
	)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error make query %s : %s", queueName, err))
	}

	return queue, nil
}

func MakeConsumer(ch wabbit.Channel, queueName string, currentQueue ...wabbit.Queue) (<-chan wabbit.Delivery, error) {
	var err error
	var queue wabbit.Queue
	if len(currentQueue) > 0 {
		queue = currentQueue[0]
	} else {
		queue, err = MakeQueue(ch, queueName)

		if err != nil {
			return nil, err
		}
	}

	consumer, err := ch.Consume(queue.Name(), "", nil)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error build consumer for queue %s : %s", queue.Name(), err))
	}

	return consumer, nil
}

package main

import (
	"JMVC-Middleware/config"
	"JMVC-Middleware/connection"
	"JMVC-Middleware/controller"
	utilities "JMVC-Middleware/utility"
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/streadway/amqp"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			stack := utilities.Stack(3)
			utilities.Error.Fatalf("[Recovery] %s panic recovered:\n%s\n%s%s", time.Now().Format("2006-01-02 15:04:05"), err, stack, "\033[0m")
		}
	}()

	// ch, err := RabbitInit()

	// //Retry the connection 6 times for every 10 seconds if the connection is error and always retrying if IS_PRODUCTION is true
	// for i := 1; (i <= 6 || config.IS_PRODUCTION) && err != nil; i++ {
	// 	time.Sleep(time.Second * 10)
	// 	ch, err = RabbitInit()
	// }

	// if err != nil {
	// 	utilities.Error.Fatal(err)
	// }

	// defer ch.Close()

	cInflux, err := InfluxInit()

	//Retry the connection 6 times for every 10 seconds if the connection is error and always retrying if IS_PRODUCTION is true
	for i := 1; (i <= 3 || config.IS_PRODUCTION) && err != nil; i++ {
		time.Sleep(time.Second * 10)
		cInflux, err = InfluxInit()
	}

	if err != nil {
		utilities.Error.Fatal(err.Error())
	}

	// counterConsumer, err := connection.MakeConsumer(ch, config.RABBIT_QUEUE)

	// if err != nil {
	// 	utilities.Error.Fatal(err)
	// }
	amqpServerURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RABBIT_USERNAME, config.RABBIT_PASSWORD, config.RABBIT_ADDRESS, config.RABBIT_PORT)
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		utilities.Error.Println(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		utilities.Error.Println(err)
	}
	defer channelRabbitMQ.Close()

	// Subscribing to QueueService1 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		config.RABBIT_QUEUE, // queue name
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no local
		false,               // no wait
		nil,                 // arguments
	)
	if err != nil {
		utilities.Error.Println(err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range messages {
			_ = controller.InsertCounterData(msg, cInflux)
		}
	}()

	fmt.Println("Waiting Sensor Data...")
	<-forever

	utilities.Error.Fatalln("Unexpected Shutdown")
}

//Start connection to InfluxDB
func InfluxInit() (client.Client, error) {
	//Start dialing to InfluxDB
	cInflux, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.InfluxAddress + ":" + config.InfluxPort,
		Username: config.InfluxUsername,
		Password: config.InfluxPassword,
	})

	if err == nil {
		_, _, err = cInflux.Ping(time.Second * 3)
	}

	if err != nil {
		return nil, fmt.Errorf("error connect to influxdb : %s", err)
	}

	//Create database based on InfluxDB configuration
	_, err = connection.QueryDB(fmt.Sprintf("CREATE DATABASE %s", config.InfluxDatabaseName), cInflux)

	if err != nil {
		return nil, fmt.Errorf("error creating database influxdb : %s", err)
	}

	return cInflux, nil
}

// //Connect to RabbitMQ message broker
// func RabbitInit() (wabbit.Channel, error) {
// 	//Set connection string that contains username, password, address, and port for connecting to rabbitMQ
// 	conString := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RABBIT_USERNAME, config.RABBIT_PASSWORD, config.RABBIT_ADDRESS, config.RABBIT_PORT)

// 	conn, err := amqp.Dial(conString)

// 	var ch wabbit.Channel

// 	if err == nil {
// 		ch, err = conn.Channel()
// 	}

// 	return ch, err
// }

package main

import (
	"reflect"

	"github.com/caarlos0/env"
	gnats "github.com/getmilly/grok/nats"
	"github.com/getmilly/levels/pkg/messages"
	"github.com/joho/godotenv"
	stan "github.com/nats-io/go-nats-streaming"
	uuid "github.com/satori/go.uuid"
)

type configuration struct {
	Queue       string `env:"QUEUE"`
	Subject     string `env:"SUBJECT"`
	NatsURL     string `env:"NATS_URL"`
	ClusterName string `env:"CLUSTER_NAME"`
}

func main() {
	conf := readConfiguration()

	conn, err := stan.Connect(
		conf.ClusterName,
		uuid.NewV4().String(),
		stan.NatsURL(conf.NatsURL),
	)

	if err != nil {
		panic(err)
	}

	gnats.NewSubscriber(conn).
		WithQueue(conf.Queue).
		WithSubject(conf.Subject).
		WithMessageType(reflect.TypeOf(messages.RewardReached{})).
		Run()
}

func readConfiguration() *configuration {
	conf := new(configuration)

	godotenv.Load()
	err := env.Parse(conf)

	if err != nil {
		panic(err)
	}

	return conf
}

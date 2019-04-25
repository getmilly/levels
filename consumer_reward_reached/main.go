package main

import (
	"reflect"

	"github.com/getmilly/grok/mongodb"

	"github.com/getmilly/levels/pkg/level"
	"github.com/getmilly/levels/pkg/repositories"

	"github.com/getmilly/levels/pkg/services"

	"github.com/caarlos0/env"
	gnats "github.com/getmilly/grok/nats"
	"github.com/getmilly/levels/pkg/messages"
	"github.com/joho/godotenv"
	stan "github.com/nats-io/go-nats-streaming"
	uuid "github.com/satori/go.uuid"
)

//Settings wraps all application settingsigs needs
type Settings struct {
	Queue                string `env:"QUEUE,required"`
	Subject              string `env:"SUBJECT,required"`
	NatsURL              string `env:"NATS_URL,required"`
	ClusterName          string `env:"CLUSTER_NAME,required"`
	MongoURL             string `env:"MONGO_URL,required"`
	LevelUpdatedSubject  string `env:"LEVEL_UPDATED_SUBJECT,required"`
	LevelUpgradedSubject string `env:"LEVEL_UPGRADED_SUBJECT,required"`
}

func main() {
	settings := readSettings()

	conn, err := stan.Connect(
		settings.ClusterName,
		uuid.NewV4().String(),
		stan.NatsURL(settings.NatsURL),
	)

	if err != nil {
		panic(err)
	}

	client, err := mongodb.Connect(settings.MongoURL)

	if err != nil {
		panic(err)
	}

	levelService := services.NewLevelService(
		settings.LevelUpgradedSubject,
		settings.LevelUpdatedSubject,
		gnats.NewProducer(conn),
		level.NewTibiaCalculator(),
		repositories.NewExperienceHistory(client),
		repositories.NewCalculatedLevelRepository(client),
	)

	gnats.NewSubscriber(conn).
		WithQueue(settings.Queue).
		WithSubject(settings.Subject).
		WithMessageType(reflect.TypeOf(messages.RewardReached{})).
		WithHandler(func(message interface{}) error {
			return levelService.HandleReward(message.(*messages.RewardReached))
		}).
		Run()
}

func readSettings() *Settings {
	settings := new(Settings)

	godotenv.Load()
	err := env.Parse(settings)

	if err != nil {
		panic(err)
	}

	return settings
}

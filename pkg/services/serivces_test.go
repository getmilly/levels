package services_test

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type ServiceSettings struct {
	NatsURL     string `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	NatsCluster string `env:"NATS_CLUSTER" envDefault:"test-cluster"`
	MongoURL    string `env:"MONGO_URL" envDefault:"mongodb://localhost:27017"`
}

func LoadSettings() *ServiceSettings {
	settings := new(ServiceSettings)

	godotenv.Load()
	err := env.Parse(settings)

	if err != nil {
		panic(err)
	}

	return settings
}

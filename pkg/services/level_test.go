package services_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/getmilly/levels/pkg/messages"
	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/mongo"

	stan "github.com/nats-io/go-nats-streaming"
	uuid "github.com/satori/go.uuid"

	"github.com/getmilly/levels/pkg/level"
	"github.com/getmilly/levels/pkg/models"
	"github.com/getmilly/levels/pkg/repositories"

	mongodb "github.com/getmilly/grok/mongodb"
	"github.com/getmilly/grok/nats"
	"github.com/getmilly/levels/pkg/services"
)

func TestLevel_HandleReward_Upgrade(t *testing.T) {
	natsConn := getNatsConn()
	mongoClient := getMongoClient()
	service := services.NewLevelService(
		"level-upgraded",
		"level-updated",
		nats.NewProducer(natsConn),
		level.NewTibiaCalculator(),
		repositories.NewExperienceHistory(mongoClient),
		repositories.NewCalculatedLevelRepository(mongoClient),
	)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		nats.NewSubscriber(natsConn).
			WithQueue(uuid.NewV4().String()).
			WithSubject("level-upgraded").
			WithMessageType(reflect.TypeOf(models.CalculatedLevel{})).
			WithHandler(func(m interface{}) error {
				defer wg.Done()
				return nil
			}).
			Run()
	}()

	playerID := uuid.NewV4().String()

	err := service.HandleReward(&messages.RewardReached{
		Kind:     "xp",
		PlayerID: playerID,
		Data: &messages.RewardExperience{
			Experience: 10,
		},
	})

	err = service.HandleReward(&messages.RewardReached{
		Kind:     "xp",
		PlayerID: playerID,
		Data: &messages.RewardExperience{
			Experience: 100,
		},
	})

	assert.NoError(t, err)

	wg.Wait()
}

func TestLevel_HandleReward_Update(t *testing.T) {
	natsConn := getNatsConn()
	mongoClient := getMongoClient()
	service := services.NewLevelService(
		"level-upgraded",
		"level-updated",
		nats.NewProducer(natsConn),
		level.NewTibiaCalculator(),
		repositories.NewExperienceHistory(mongoClient),
		repositories.NewCalculatedLevelRepository(mongoClient),
	)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		nats.NewSubscriber(natsConn).
			WithQueue(uuid.NewV4().String()).
			WithSubject("level-updated").
			WithMessageType(reflect.TypeOf(models.CalculatedLevel{})).
			WithHandler(func(m interface{}) error {
				defer wg.Done()
				return nil
			}).
			Run()
	}()

	playerID := uuid.NewV4().String()

	err := service.HandleReward(&messages.RewardReached{
		Kind:     "xp",
		PlayerID: playerID,
		Data: &messages.RewardExperience{
			Experience: 10,
		},
	})

	err = service.HandleReward(&messages.RewardReached{
		Kind:     "xp",
		PlayerID: playerID,
		Data: &messages.RewardExperience{
			Experience: 10,
		},
	})

	assert.NoError(t, err)

	wg.Wait()
}

func getNatsConn() stan.Conn {
	settings := LoadSettings()
	conn, err := stan.Connect(settings.NatsCluster, uuid.NewV4().String(), stan.NatsURL(settings.NatsURL))

	if err != nil {
		panic(err)
	}

	return conn
}

func getMongoClient() *mongo.Client {
	settings := LoadSettings()
	client, err := mongodb.Connect(settings.MongoURL)

	if err != nil {
		panic(err)
	}

	return client
}

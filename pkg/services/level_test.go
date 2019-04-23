package services_test

import (
	"os"
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

func TestLevel_HandleReward(t *testing.T) {
	natsConn := getNatsConn()
	service := services.NewLevelService(
		"level-upgraded",
		"level-updated",
		nats.NewProducer(natsConn),
		level.NewTibiaCalculator(),
		repositories.NewCalculatedLevelRepository(getMongoClient()),
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

	err := service.HandleReward(&messages.RewardReached{
		PlayerID: uuid.NewV4().String(),
		Data: &messages.RewardExperience{
			Experience: 10,
		},
	})

	err = service.HandleReward(&messages.RewardReached{
		PlayerID: uuid.NewV4().String(),
		Data: &messages.RewardExperience{
			Experience: 100,
		},
	})

	assert.NoError(t, err)

	wg.Wait()
}

func getNatsConn() stan.Conn {
	natsURL := os.Getenv("NATS_URL")
	clusterID := os.Getenv("NATS_CLUSTER")

	conn, err := stan.Connect(clusterID, uuid.NewV4().String(), stan.NatsURL(natsURL))

	if err != nil {
		panic(err)
	}

	return conn
}

func getMongoClient() *mongo.Client {
	client, err := mongodb.Connect(os.Getenv("MONGO_URL"))

	if err != nil {
		panic(err)
	}

	return client
}

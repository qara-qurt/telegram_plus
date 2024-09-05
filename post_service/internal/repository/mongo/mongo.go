package mongo

import (
	"context"
	"fmt"

	"github.com/qara-qurt/telegram_plus/post_service/internal/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Mongodb *mongo.Database
}

func NewMongo(conf config.MongoDB) (*Mongo, error) {

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DBName, conf.AuthSource)

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Errorf("error while connecting to the database, %v", err)
		return nil, fmt.Errorf("error while connecting to the database, %v", err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		logrus.Errorf("error while pinging the database, %v", err)
		return nil, fmt.Errorf("error while pinging the database, %v", err)
	}

	database := client.Database("post_service")

	logrus.Info("Connected to MongoDB!")

	return &Mongo{
		Mongodb: database,
	}, nil
}

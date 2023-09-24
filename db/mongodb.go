package db

import (
	"context"
	"fmt"

	"github.com/fhy/utils-golang/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongodbInit(_config *config.MongoDBConfig) (*mongo.Database, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/?timeoutMS=%d", _config.Username, _config.Password, _config.Host, _config.Port, _config.Timeout)
	loggerOptions := options.
		Logger().
		SetComponentLevel(options.LogComponentCommand, _config.LogLevel)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).SetLoggerOptions(loggerOptions)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		logrus.Panicf("error initing mongodb err: %s", err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			logrus.Panicf("error initing mongodb err: %s", err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		logrus.Panicf("error initing mongodb err: %s", err)
	}
	logrus.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client.Database(_config.DB), nil
}

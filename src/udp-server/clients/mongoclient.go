package clients

import (
	"context"
	"fmt"
	"main/helpers"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func (client MongoClient) Disconnect(ctx context.Context) error {
	logrus.Warningln("Closing MongoDB connection!")
	return client.Client.Disconnect(ctx)
}

func (client MongoClient) Insert(ctx context.Context, packetType string, packet interface{}) error {
	logrus.Infoln("Inserted Document to MongoDB!")
	collection := client.Database.Collection(packetType)
	if collection == nil {
		return fmt.Errorf("Collection could not been retrieved")
	}

	result, err := collection.InsertOne(ctx, packet)
	helpers.LogIfError(err)

	logrus.Debugln("Inserted a single document: ", result.InsertedID)

	return err
}

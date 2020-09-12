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

func (client MongoClient) Connect(ctx context.Context) error {
	logrus.Infoln("Connected to MongoDB!")

	return fmt.Errorf("NOT implemented for MongoDB!")
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

func (client MongoClient) Update(packet interface{}) error {
	logrus.Infoln("Updating Document in MongoDB!")
	return fmt.Errorf("NOT implemented for MongoDB!")
}

func (client MongoClient) Delete(id string) error {
	info := fmt.Sprintf("Removing Document with id: %s from MongoDB!", id)
	logrus.Infoln(info)
	return fmt.Errorf("NOT implemented for MongoDB!")
}

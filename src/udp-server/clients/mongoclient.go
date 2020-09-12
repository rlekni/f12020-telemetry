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

func (client MongoClient) InsertPacketMotionData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketMotionData, packet)
}

func (client MongoClient) InsertPacketSessionData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketSessionData, packet)
}

func (client MongoClient) InsertPacketLapData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketLapData, packet)
}

func (client MongoClient) InsertPacketEventData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketEventData, packet)
}

func (client MongoClient) InsertPacketParticipantsData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketParticipantsData, packet)
}

func (client MongoClient) InsertPacketCarSetupData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketCarSetupData, packet)
}

func (client MongoClient) InsertPacketCarTelemetryData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketCarTelemetryData, packet)
}

func (client MongoClient) InsertPacketCarStatusData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketCarStatusData, packet)
}

func (client MongoClient) InsertPacketFinalClassificationData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketFinalClassificationData, packet)
}

func (client MongoClient) InsertPacketLobbyInfoData(ctx context.Context, packet interface{}) error {
	return client.Insert(ctx, PacketLobbyInfoData, packet)
}

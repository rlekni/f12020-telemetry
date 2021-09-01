package clients

import (
	"context"
	"main/helpers"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RepositoryClient -
type RepositoryClient interface {
	Disconnect(ctx context.Context) error
	Insert(ctx context.Context, packetType string, packet interface{}) error
	InsertPacketMotionData(ctx context.Context, packet interface{}) error
	InsertPacketSessionData(ctx context.Context, packet interface{}) error
	InsertPacketLapData(ctx context.Context, packet interface{}) error
	InsertPacketEventData(ctx context.Context, packet interface{}) error
	InsertPacketParticipantsData(ctx context.Context, packet interface{}) error
	InsertPacketCarSetupData(ctx context.Context, packet interface{}) error
	InsertPacketCarTelemetryData(ctx context.Context, packet interface{}) error
	InsertPacketCarStatusData(ctx context.Context, packet interface{}) error
	InsertPacketFinalClassificationData(ctx context.Context, packet interface{}) error
	InsertPacketLobbyInfoData(ctx context.Context, packet interface{}) error
}

func NewRepositoryClient(ctx context.Context) RepositoryClient {
	logrus.Debugln("Creating MongoDB client.")

	connectionString := os.Getenv("MONGO_CONNECTION_STRING")
	databaseName := os.Getenv("MONGO_DATABASE")
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	helpers.ThrowIfError(err)

	// Check the connection
	err = client.Ping(ctx, nil)
	helpers.ThrowIfError(err)

	logrus.Infoln("Connected to MongoDB!")
	database := client.Database(databaseName)

	return &MongoClient{
		Client:   client,
		Database: database,
	}
}

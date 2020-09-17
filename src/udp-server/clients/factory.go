package clients

import (
	"context"
	"database/sql"
	"fmt"
	"main/helpers"
	"os"
	"strconv"

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

type RepositoryType string

const (
	Mongo   RepositoryType = "MONGODB"
	Postgre RepositoryType = "POSTGRES"
)

func NewRepositoryClient(ctx context.Context, repositoryType RepositoryType) RepositoryClient {
	switch repositoryType {
	case Mongo:
		return newMongoClient(ctx)
	case Postgre:
		return newPostgreClient()
	default:
		warning := fmt.Sprintf("No repository type matched. Provided: %s", repositoryType)
		logrus.Warningln(warning)
		return nil
	}
}

func newMongoClient(ctx context.Context) *MongoClient {
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

func newPostgreClient() *PostgreClient {
	logrus.Debugln("Creating PostgreSQL client.")

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")
	port, err := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 64)
	helpers.ThrowIfError(err)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	helpers.ThrowIfError(err)

	err = db.Ping()
	helpers.ThrowIfError(err)

	logrus.Infoln("Connected to PostgreSQL!")
	return &PostgreClient{
		Database: db,
	}
}

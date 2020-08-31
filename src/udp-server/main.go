package main

import (
	"context"
	"main/f12020packets"
	"net"
	"os"
	"time"

	"github.com/snowzach/rotatefilehook"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	initialiseLogger()

	listenPort := os.Getenv("UDP_PORT")
	logrus.Infoln("Using Port: ", listenPort)

	PORT := ":" + listenPort
	connection, err := initialiseUDPConnection(PORT)
	if err != nil {
		logrus.Fatalln(err)
		return
	}
	defer connection.Close()

	mongoClient, ctx := initialiseMongoDBConnection()
	defer mongoClient.Disconnect(ctx)

	mongoDatabase := mongoClient.Database("f12020telemetry")

	buffer := make([]byte, 2048)
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		logrus.Info("-> ", addr)
		data := buffer[0:n]

		deserialisePacket(ctx, mongoDatabase, data)
	}
}

func initialiseLogger() {
	var logLevel = logrus.InfoLevel

	logFileDirectory := os.Getenv("LOGS_PATH")
	filename := logFileDirectory + "/console.log"
	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   filename,
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     28,
		Level:      logLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	logrus.SetLevel(logLevel)
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	logrus.AddHook(rotateFileHook)
}

func initialiseUDPConnection(port string) (*net.UDPConn, error) {
	s, err := net.ResolveUDPAddr("udp4", "0.0.0.0"+port)
	if err != nil {
		return nil, err
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func initialiseMongoDBConnection() (*mongo.Client, context.Context) {
	ctx := context.Background()

	connectionString := os.Getenv("MONGO_CONNECTION_STRING")

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		logrus.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infoln("Connected to MongoDB!")
	return client, ctx
}

func deserialisePacket(ctx context.Context, mongoDatabase *mongo.Database, data []byte) {
	header, err := f12020packets.ToPacketHeader(data[0:24])
	if err != nil {
		logrus.Errorf("Failed to decode Packet Header. Error: %q", err)
	}

	logrus.Debugln("SessionID: ", header.SessionUID)
	logrus.Debugf("Data length: %d, PacketID: %d\n", len(data), header.PacketID)

	switch header.PacketID {
	case 0:
		packet, err := f12020packets.ToPacketMotionData(data[24:1464], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetMotionData", packet)
	case 1:
		packet, err := f12020packets.ToPacketSessionData(data[24:251], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetSessionData", packet)
	case 2:
		packet, err := f12020packets.ToPacketLapData(data[24:1190], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetLapData", packet)
	case 3:
		packet, err := f12020packets.ToPacketEventData(data[24:35], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetEventData", packet)
	case 4:
		packet, err := f12020packets.ToPacketParticipantsData(data[24:1213], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetParticipantsData", packet)
	case 5:
		packet, err := f12020packets.ToPacketCarSetupData(data[24:1102], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetCarSetupData", packet)
	case 6:
		packet, err := f12020packets.ToPacketCarTelemetryData(data[24:1307], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetCarTelemetryData", packet)
	case 7:
		packet, err := f12020packets.ToPacketCarStatusData(data[24:1344], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetCarStatusData", packet)
	case 8:
		packet, err := f12020packets.ToPacketFinalClassificationData(data[24:839], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetFinalClassificationData", packet)
	case 9:
		packet, err := f12020packets.ToPacketLobbyInfoData(data[24:1169], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetLobbyInfoData", packet)
	default:
		logrus.Warningf("None of the defined PacketIDs matched. Data length: %d, PacketID: %d\n", len(data), header.PacketID)
	}
	// insertDocument(ctx, mongoDatabase, "packer", packet)
}

func insertDocument(ctx context.Context, mongoDatabase *mongo.Database, collectionName string, packet interface{}) {
	collection := mongoDatabase.Collection(collectionName)
	insertResult, err := collection.InsertOne(ctx, packet)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Debugln("Inserted a single document: ", insertResult.InsertedID)
}

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
	initLogger()
	arguments := os.Args
	portInput := "20777"
	if len(arguments) == 1 {
		logrus.Infoln("NOTE: port number not provided, defaulting to 20777")
	} else {
		portInput = arguments[1]
	}
	PORT := ":" + portInput

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	defer connection.Close()

	// Make buffer big enough 2048 should be enough
	buffer := make([]byte, 2048)
	mongoClient, ctx := newMongoDBConnection()
	mongoDatabase := mongoClient.Database("f12020telemetry")
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

func deserialisePacket(ctx context.Context, mongoDatabase *mongo.Database, data []byte) {
	header, err := f12020packets.ToPacketHeader(data[0:24])
	if err != nil {
		logrus.Errorf("Failed to decode Packet Header. Error: %q", err)
	}

	logrus.Debugln("SessionID: ", header.SessionUID)
	logrus.Debugf("Data length: %d, PacketID: %d\n", len(data), header.PacketID)
	switch header.PacketID {
	case 0:
		logrus.Debugln("Decoding PacketMotionData")
		packet, err := f12020packets.ToPacketMotionData(data[24:1464], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetMotionData", packet)
	case 1:
		logrus.Debugln("Decoding PacketSessionData")
		packet, err := f12020packets.ToPacketSessionData(data[24:251], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetSessionData", packet)
	case 2:
		logrus.Debugln("Decoding PacketLapData")
		packet, err := f12020packets.ToPacketLapData(data[24:1190], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetLapData", packet)
	case 3:
		logrus.Debugln("Decoding PacketEventData")
		packet, err := f12020packets.ToPacketEventData(data[24:35], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetEventData", packet)
	case 4:
		logrus.Debugln("Decoding PacketParticipantsData")
		packet, err := f12020packets.ToPacketParticipantsData(data[24:1213], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetParticipantsData", packet)
	case 5:
		logrus.Debugln("Decoding PacketCarSetupData")
		packet, err := f12020packets.ToPacketCarSetupData(data[24:1102], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetCarSetupData", packet)
	case 6:
		logrus.Debugln("Decoding PacketCarTelemetryData")
		packet, err := f12020packets.ToPacketCarTelemetryData(data[24:1307], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetCarTelemetryData", packet)
	case 7:
		logrus.Debugln("Decoding PacketCarStatusData")
		packet, err := f12020packets.ToPacketCarStatusData(data[24:1344], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetCarStatusData", packet)
	case 8:
		logrus.Debugln("Decoding PacketFinalClassificationData")
		packet, err := f12020packets.ToPacketFinalClassificationData(data[24:839], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetFinalClassificationData", packet)
	case 9:
		logrus.Debugln("Decoding PacketLobbyInfoData")
		packet, err := f12020packets.ToPacketLobbyInfoData(data[24:1169], header)
		if err != nil {
			logrus.Errorln(err)
		}
		insertDocument(ctx, mongoDatabase, "packetLobbyInfoData", packet)
	default:
		logrus.Warningf("None of the defined PacketIDs matched. Data length: %d, PacketID: %d\n", len(data), header.PacketID)
	}
}

func insertDocument(ctx context.Context, mongoDatabase *mongo.Database, collectionName string, packet interface{}) {
	collection := mongoDatabase.Collection(collectionName)
	insertResult, err := collection.InsertOne(ctx, packet)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Debugln("Inserted a single document: ", insertResult.InsertedID)
}

func newMongoDBConnection() (*mongo.Client, context.Context) {
	ctx := context.Background()

	// username := "telemetry_user"
	// password := ""
	// connectionString := fmt.Sprintf("mongodb+srv://%s:%s@test.azure.mongodb.net/f12020telemetry?retryWrites=true&w=majority", username, password)
	connectionString := "mongodb://localhost:27017/f12020telemetry?retryWrites=true&w=majority"
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

// Solution from https://github.com/sirupsen/logrus/issues/784#issuecomment-403765306
func initLogger() {
	var logLevel = logrus.InfoLevel

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logs/console.log",
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

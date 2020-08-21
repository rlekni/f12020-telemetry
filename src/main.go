package main

import (
	"context"
	"fmt"
	"log"
	"main/f12020packets"
	"net"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	arguments := os.Args
	portInput := "20777"
	if len(arguments) == 1 {
		fmt.Println("NOTE: port number not provided, defaulting to 20777")
	} else {
		portInput = arguments[1]
	}
	PORT := ":" + portInput

	// connection, err := net.ListenPacket("udp4", ":20777")
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
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
			fmt.Println(err)
			// Log
			continue
		}
		fmt.Print("-> ", addr)
		data := buffer[0:n]

		deserialisePacket(ctx, mongoDatabase, data)
	}
}

func deserialisePacket(ctx context.Context, mongoDatabase *mongo.Database, data []byte) {
	header, err := f12020packets.ToPacketHeader(data[0:24])
	if err != nil {
		fmt.Printf("Failed to deserialise to Packet Header")
	}

	fmt.Printf(" Data length: %d, PacketID: %d\n", len(data), header.PacketID)
	switch header.PacketID {
	case 0:
		fmt.Print("Deserialising PacketMotionData")
		packet, _ := f12020packets.ToPacketMotionData(data[24:1464], header)
		// json, _ := json.Marshal(packet)
		// fmt.Println(string(json))
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetMotionData", packet)
		// collection := mongoDatabase.Collection("packetMotionData")

		// insertResult, err := collection.InsertOne(ctx, packet)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	case 1:
		fmt.Print("Deserialising PacketSessionData")
		packet, _ := f12020packets.ToPacketSessionData(data[24:251], header)
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetMotionData", packet)
	case 2:
		fmt.Print("Deserialising PacketLapData")
		packet, _ := f12020packets.ToPacketLapData(data[24:1190], header)
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetLapData", packet)
	case 3:
		fmt.Print("Deserialising PacketEventData")
		packet, _ := f12020packets.ToPacketEventData(data[24:35], header)
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetEventData", packet)
	case 4:
		fmt.Print("Deserialising PacketParticipantsData")
		packet, _ := f12020packets.ToPacketParticipantsData(data[24:1213], header)
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetParticipantsData", packet)
	case 5:
		fmt.Print("Deserialising PacketCarSetupData")
		packet, _ := f12020packets.ToPacketCarSetupData(data[24:1102], header)
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetCarSetupData", packet)
	case 6:
		fmt.Print("Deserialising PacketCarTelemetryData")
		packet, _ := f12020packets.ToPacketCarTelemetryData(data[24:1307], header)
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetCarTelemetryData", packet)
	case 7:
		fmt.Print("Deserialising PacketCarStatusData")
		packet, _ := f12020packets.ToPacketCarStatusData(data[24:1344], header)
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetCarStatusData", packet)
	case 8:
		fmt.Print("Deserialising PacketFinalClassificationData")
		packet, _ := f12020packets.ToPacketFinalClassificationData(data[24:839], header)
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetFinalClassificationData", packet)
	case 9:
		fmt.Print("Deserialising PacketLobbyInfoData")
		packet, _ := f12020packets.ToPacketLobbyInfoData(data[24:1169], header)
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		insertDocument(ctx, mongoDatabase, "packetLobbyInfoData", packet)
	default:
		fmt.Printf("None of the defined PacketIDs matched. Data length: %d, PacketID: %d\n", len(data), header.PacketID)
	}
}

func insertDocument(ctx context.Context, mongoDatabase *mongo.Database, collectionName string, packet interface{}) {
	collection := mongoDatabase.Collection(collectionName)

	insertResult, err := collection.InsertOne(ctx, packet)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func newMongoDBConnection() (*mongo.Client, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@planner-core-test-free.tlrom.azure.mongodb.net?retryWrites=true&w=majority", username, password)
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	client, err := mongo.Connect(ctx, clientOptions)

	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx
}

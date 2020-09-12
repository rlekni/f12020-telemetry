package main

import (
	"context"
	"main/clients"
	"main/internal"

	"github.com/sirupsen/logrus"
)

func main() {
	internal.InitialiseLogger()

	// Initialise UDP listener and defer connection close
	udpConnection := internal.InitialiseUDPListener()
	defer udpConnection.Close()

	// Initialise new MongoDB connection and defer close
	ctx := context.Background()
	mongoClient := clients.NewMongoDBConnection(ctx)
	defer mongoClient.Disconnect(ctx)

	mongoClient.GetDatabase("f12020telemetry")

	buffer := make([]byte, 2048)
	for {
		n, addr, err := udpConnection.ReadFromUDP(buffer)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		logrus.Info("-> ", addr)
		data := buffer[0:n]

		internal.DeserialisePacket(ctx, mongoClient, data)
	}
}

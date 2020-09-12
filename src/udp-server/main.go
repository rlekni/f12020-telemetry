package main

import (
	"context"
	"main/clients"
	"main/internal"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	internal.InitialiseLogger()

	// Initialise UDP listener and defer connection close
	udpConnection := internal.InitialiseUDPListener()
	defer udpConnection.Close()

	// Initialise new MongoDB connection and defer close
	ctx := context.Background()
	repositoryType := os.Getenv("REPOSITORY_TYPE")
	client := clients.NewRepositoryClient(ctx, clients.RepositoryType(repositoryType))
	defer client.Disconnect(ctx)

	buffer := make([]byte, 2048)
	for {
		n, addr, err := udpConnection.ReadFromUDP(buffer)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		logrus.Info("-> ", addr)
		data := buffer[0:n]

		internal.DeserialisePacket(ctx, client, data)
	}
}

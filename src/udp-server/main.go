package main

import (
	"main/internal"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	internal.InitialiseLogger()

	// Initialise UDP listener and defer connection close
	udpConnection := internal.InitialiseUDPListener()
	defer udpConnection.Close()

	buffer := make([]byte, 2048)
	for {
		n, addr, err := udpConnection.ReadFromUDP(buffer)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		logrus.Info("-> ", addr)
		data := buffer[0:n]

		internal.DeserialisePacket(data)
	}
}

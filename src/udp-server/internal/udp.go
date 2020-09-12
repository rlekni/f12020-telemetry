package internal

import (
	"net"
	"os"

	"github.com/sirupsen/logrus"
)

func InitialiseUDPListener() *net.UDPConn {
	port := os.Getenv("UDP_PORT")
	logrus.Infoln("Using Port: ", port)

	s, err := net.ResolveUDPAddr("udp4", "0.0.0.0:"+port)
	if err != nil {
		logrus.Fatalln(err)
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		logrus.Fatalln(err)
	}

	return connection
}

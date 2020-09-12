package internal

import (
	"main/helpers"
	"net"
	"os"

	"github.com/sirupsen/logrus"
)

func InitialiseUDPListener() *net.UDPConn {
	port := os.Getenv("UDP_PORT")
	logrus.Infoln("Using Port: ", port)

	s, err := net.ResolveUDPAddr("udp4", "0.0.0.0:"+port)
	helpers.ThrowIfError(err)

	connection, err := net.ListenUDP("udp4", s)
	helpers.ThrowIfError(err)

	return connection
}

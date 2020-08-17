package main

import (
	"encoding/gob"
	"fmt"
	"main/f12020packets"
	"math/rand"
	"net"
	"os"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

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

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			// Log
			continue
		}
		fmt.Print("-> ", addr)
		data := buffer[0:n]
		fmt.Printf(" Data length: %d", len(data))

		// var network bytes.Buffer
		// network.Write(data)
		enc := gob.NewDecoder(connection)
		var packet f12020packets.TestDataPacket
		err = enc.Decode(&packet)
		if err != nil {
			fmt.Println(err)
			// Log
			continue
		}
		fmt.Printf(" Packet: %d, %d.%d\n", packet.PacketFormat, packet.GameMajorVersion, packet.GameMinorVersion)
		// return
	}
}

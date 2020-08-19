package main

import (
	"fmt"
	"main/f12020packets"
	"net"
	"os"
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

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			// Log
			continue
		}
		fmt.Print("-> ", addr)
		data := buffer[0:n]
		fmt.Printf(" Data length: %d, byte 0: %q\n", len(data), uint8(data[2]))

		if len(data) == 1464 {
			packet, err := f12020packets.ToPacketMotionData(data[0:1464])
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		}
	}
}

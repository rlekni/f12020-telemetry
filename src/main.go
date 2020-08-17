package main

import (
	"encoding/binary"
	"fmt"
	"main/f12020packets"
	"math"
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
		fmt.Printf(" Data length: %d, byte 0: %q\n", len(data), uint8(data[2]))

		if len(data) == 1464 {
			// enc := gob.NewDecoder(connection)
			// var packet f12020packets.PacketMotionData
			// err = enc.Decode(&packet)

			packet := f12020packets.PacketMotionData{
				Header: f12020packets.PacketHeader{
					PacketFormat:            binary.LittleEndian.Uint16(data[0:2]),
					GameMajorVersion:        uint8(data[2]),
					GameMinorVersion:        uint8(data[3]),
					PacketVersion:           uint8(data[4]),
					PacketID:                uint8(data[5]),
					SessionUID:              binary.LittleEndian.Uint64(data[6:14]),
					SessionTime:             math.Float32frombits(binary.LittleEndian.Uint32(data[14:18])),
					FrameIdentifier:         binary.LittleEndian.Uint32(data[18:22]),
					PlayerCarIndex:          uint8(22),
					SecondaryPlayerCarIndex: uint8(23),
				},
			}
			fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
		}
		if err != nil {
			fmt.Println(err)
			// Log
			continue
		}
	}
}

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

		deserialisePacket(data)
	}
}

func deserialisePacket(data []byte) {
	fmt.Printf(" Data length: %d\n", len(data))
	switch len(data) {
	case 1464:
		fmt.Print("Deserialising PacketMotionData")
		packet, _ := f12020packets.ToPacketMotionData(data[0:1464])
		// json, _ := json.Marshal(packet)
		// fmt.Println(string(json))
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	case 251:
		fmt.Print("Deserialising PacketSessionData")
		packet, _ := f12020packets.ToPacketSessionData(data[0:251])
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	case 1190:
		fmt.Print("Deserialising PacketLapData")
		packet, _ := f12020packets.ToPacketLapData(data[0:1190])
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	case 35:
		fmt.Print("Deserialising PacketEventData")
		packet, _ := f12020packets.ToPacketEventData(data[0:35])
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	case 1213:
		fmt.Print("Deserialising PacketParticipantsData")
		packet, _ := f12020packets.ToPacketParticipantsData(data[0:1213])
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	case 1102:
		fmt.Print("Deserialising PacketCarSetupData")
		packet, _ := f12020packets.ToPacketCarSetupData(data[0:1102])
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	case 1307:
		fmt.Print("Deserialising PacketCarTelemetryData")
		packet, _ := f12020packets.ToPacketCarTelemetryData(data[0:1307])
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	case 1344:
		fmt.Print("Deserialising PacketCarStatusData")
		packet, _ := f12020packets.ToPacketCarStatusData(data[0:1344])
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	case 839:
		fmt.Print("Deserialising PacketFinalClassificationData")
		packet, _ := f12020packets.ToPacketFinalClassificationData(data[0:839])
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	case 1169:
		fmt.Print("Deserialising PacketLobbyInfoData")
		packet, _ := f12020packets.ToPacketLobbyInfoData(data[0:1169])
		fmt.Printf(" Packet: %d, %d.%d\n", packet.Header.PacketFormat, packet.Header.GameMajorVersion, packet.Header.GameMinorVersion)
	default:
		fmt.Printf("None of the defined lengths matched. Data length: %d", len(data))
	}
}

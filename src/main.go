package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
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
		fmt.Print("-> ", addr)

		// Stop the server if STOP command is received
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := buffer[0:n]
		fmt.Printf(" Data length: %d \n", len(data))
		if err != nil {
			fmt.Println(err)
			// Log
			continue
		}
	}
}

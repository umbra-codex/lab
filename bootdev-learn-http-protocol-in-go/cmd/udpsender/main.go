package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const serverAddr = "localhost:42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		log.Fatalf("error resolving UDP address: %s\n", err.Error())
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("error dialing UDP connection: %s\n", err.Error())
	}
	defer conn.Close()

	fmt.Printf("Sending to %s. Type your message and press Enter to send. Press Ctrl+C to exit.\n", serverAddr)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading input: %s\n", err.Error())
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Fatalf("error sending message: %s\n", err.Error())
		}

		fmt.Printf("Message send: %s:", message)
	}
}

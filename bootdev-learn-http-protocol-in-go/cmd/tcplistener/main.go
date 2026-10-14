package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Printf("Accepted connection from %s\n\n", conn.RemoteAddr())

		linesChan := getLinesChannel(conn)
		for line := range linesChan {
			fmt.Println(line)
		}
		fmt.Println("\nConnection to ", conn.RemoteAddr(), "closed")
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {
	linesChan := make(chan string)

	go func() {
		defer file.Close()
		defer close(linesChan)

		var lines string
		for {
			buffer := make([]byte, 8, 8)
			num, err := file.Read(buffer)
			if err != nil {
				if lines != "" {
					linesChan <- lines
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}
			str := string(buffer[:num])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				linesChan <- lines + parts[i]
				lines = ""
			}
			lines += parts[len(parts)-1]
		}
	}()

	return linesChan
}

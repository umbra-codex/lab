package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const filePath = "messages.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Reading data from %s\n", filePath)
	fmt.Println("==============================")

	lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
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

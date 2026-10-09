package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const filePath = "messages.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Reading data from %s\n", filePath)
	fmt.Println("==============================")

	for {
		b := make([]byte, 8, 8)
		n, err := file.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		str := string(b[:n])
		fmt.Printf("read: %s\n", str)
	}
}

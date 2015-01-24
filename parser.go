package main

import (
	"log"
	"os"
)

func main() {
	log.Println(os.Args)
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, _ := file.Stat()
	log.Println(fileInfo.Name())
	log.Println(fileInfo.Size())
}

package main

import (
	"encoding/binary"
	"log"
	"os"
)

const MAX_OSPATH = 260

type Header struct {
	DemoFilestamp   [8]byte          // Should be HL2DEMO
	DemoProtocol    int32            // Should be DEMO_PROTOCOL
	NetworkProtocol int32            // Should be PROTOCOL_VERSION
	ServerName      [MAX_OSPATH]byte // Name of server
	ClientName      [MAX_OSPATH]byte
	MapName         [MAX_OSPATH]byte // Name of map
	GameDirectory   [MAX_OSPATH]byte // Name of game directory (com_gamedir)
	PlaybackTime    float32          // Time of track
	PlaybackTicks   int32            // # of ticks in track
	PlaybackFrames  int32            // # of frames in track
	SignonLength    int32            // length of sigondata in bytes
}

func main() {
	log.Println(os.Args)
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, _ := file.Stat()
	log.Println(fileInfo.Name())
	log.Println(fileInfo.Size())

	var header Header
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		log.Println("binary.Read failed:", err)
	}
	log.Println(string(header.ClientName[:]))
	log.Println(header.PlaybackTime)
}

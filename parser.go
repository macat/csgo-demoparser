package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

const MAX_OSPATH = 260
const DEMO_PROTOCOL = 4

type FileHeader struct {
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

type CmdHeader struct {
	Cmd        byte
	Timestamp  int32
	PlayerSlot byte
}

const (
	// it's a startup message, process as fast as possible
	DemSignon = 1
	// it's a normal network packet that we stored off
	DemPacket = 2
	// sync client clock to demo tick
	DemSynctick = 3
	// console command
	DemConsolecmd = 4
	// user input command
	DemUsercmd = 5
	// network data tables
	DemDatatables = 6
	// end of time.
	DemStop = 7
	// a blob of binary data understood by a callback function
	DemCustomdata   = 8
	DemStringtables = 9
)

func main() {
	log.Println(os.Args)
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, _ := file.Stat()
	log.Println(fileInfo.Name())
	header := parseHeader(file)
	if header.DemoProtocol != DEMO_PROTOCOL {
		log.Fatal("Bad demo protocol")
	}
	cmdHeader := parseCmdHeader(file)
	finished := false
	switch cmdHeader.Cmd {
	case DemSignon:
		log.Println("signon")
		break
	case DemUsercmd:
		log.Println("user")
		break
	case DemStop:
		finished = true
	case DemConsolecmd:
	}
	log.Println(finished)
}

func parseHeader(reader io.Reader) *FileHeader {
	header := &FileHeader{}
	err := binary.Read(reader, binary.LittleEndian, header)
	if err != nil {
		log.Println("binary.Read failed:", err)
	}
	return header
}

func parseCmdHeader(reader io.Reader) *CmdHeader {
	header := &CmdHeader{}
	err := binary.Read(reader, binary.LittleEndian, header)
	if err != nil {
		log.Println("binary.Read failed:", err)
	}
	return header
}

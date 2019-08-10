package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcapgo"

	"github.com/x-way/pktdump"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Error: missing pcap filename parameter")
	}

	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(fmt.Sprintf("Could not open pcap file '%s': %v\n", file, err))
		}
		defer f.Close()

		handle, err := pcapgo.NewReader(f)
		if err != nil {
			log.Fatal(fmt.Sprintf("Could not create pcap reader: %v\n", err))
		}

		pkgsrc := gopacket.NewPacketSource(handle, handle.LinkType())

		for packet := range pkgsrc.Packets() {
			fmt.Printf("%s ", packet.Metadata().CaptureInfo.Timestamp.Format("15:04:05.000000"))
			fmt.Println(pktdump.Format(packet))
		}
	}
}

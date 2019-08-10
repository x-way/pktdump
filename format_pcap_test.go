package pktdump_test

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcapgo"

	"github.com/x-way/pktdump"
)

func getTcpdumpOutput(filename string) string {
	out, err := exec.Command("tcpdump", "-S", "-t", "-n", "-r", filename).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func getFormatOutput(filename string) string {
	f, _ := os.Open(filename)
	defer f.Close()
	handle, err := pcapgo.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	out := ""
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		out = out + pktdump.Format(packet) + "\n"
	}
	return out
}

func TestFormatPCAP(t *testing.T) {
	files, err := ioutil.ReadDir("test")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".pcap" {
			expected := getTcpdumpOutput("test/" + f.Name())
			got := getFormatOutput("test/" + f.Name())
			if got != expected {
				t.Errorf("pcap test failed for %s, got '%s', expected '%s'", f.Name(), got, expected)
			}
		}
	}
}

package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: wol mac-address\n")
		os.Exit(1)
	}
	cnn, err := net.Dial("udp", "255.255.255.255:4000")
	if err != nil {
		log.Fatal(err)
	}
	mac, err := net.ParseMAC(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var buf = bytes.Buffer{}
	buf.Write([]byte{255, 255, 255, 255, 255, 255})
	for i := 0; i < 16; i++ {
		buf.Write([]byte(mac))
	}
	cnn.Write(buf.Bytes())
}

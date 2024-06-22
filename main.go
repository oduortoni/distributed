package main

import (
	"log"

	"github.com/oduortoni/distributed/lib/p2p/transport"
)

func main() {
	tcp := transport.NewTCP(":9000")
	err := tcp.ListenAndAccept()
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}
}

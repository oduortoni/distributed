package main

import (
	"log"

	"github.com/oduortoni/distributed/lib/p2p/tcp"
)

func main() {
	tcp := tcp.NewTCP(":9000")
	err := tcp.ListenAndAccept()
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}
	select {}
}

package transport

import (
	"fmt"
	"net"
	"sync"

	"github.com/oduortoni/distributed/lib/p2p"
)

type TCP struct {
	address  string
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]p2p.Peer
}

func NewTCP(listenAddress string) *TCP {
	return &TCP{
		address: listenAddress,
	}
}

func (tcp *TCP) ListenAndAccept() error {
	listener, err := net.Listen("tcp", tcp.address)
	if err != nil {
		return err
	}
	tcp.listener = listener

	for {
		conn, err := tcp.listener.Accept()
		if err != nil {
			fmt.Printf("ERROR: TCP accept: %q\n", err)
		} else {
			go tcp.handleConnection(conn)
		}
	}
}

func (tcp *TCP) handleConnection(conn net.Conn) {
	fmt.Printf("Incoming: %+v\n", conn)
	fmt.Fprintf(conn, "%s\n", "Hello there")
}

package tcp

import "net"

type Peer struct {
	// the peer connection
	conn net.Conn
	// if we dial and retrieve a connection, then we are sending an outbound  request
	// 	=> outbound = true
	// if we accept and retrieve a connection, then we are receining an inbound request
	// 	=> outbound = false
	outbound bool
}

func NewPeer(conn net.Conn, outbound bool) *Peer {
	return &Peer{
		conn: conn,
		outbound: outbound,
	}
}
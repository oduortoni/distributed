package p2p

// Peer is an interface that represents a remote node
type Peer interface {
	
}

// Transport is anything that handles the communication 
// between the nodes in the network.
// This can be in the form of tcp, udp, websockets, ...
type Transport interface {
	ListenAndAccept() error
}

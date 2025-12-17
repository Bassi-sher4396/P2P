package main 
import ("net"
"sync")
type Node struct {
	Id string
	Address string //ip:port
 Peers map[string]*Peer
 IdToPeerMap map[string]*net.Conn
Kademlia  []*Node
 Listener net.Listener
 Mu sync.Mutex
} 

type Peer struct {
	Id string 
	Conn net.Conn  }

	type Message struct {
		Type uint8
		Data string
	}

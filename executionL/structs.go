package main 
import ("net")
type Node struct {
	Id string
	Address string //ip:port
 Peers map[string]*Peer
 IdToPeerMap map[string]*net.Conn
 Listener net.Listener
 
} 

type Peer struct {
	Id string 
	Conn net.Conn  }

	type Message struct {
		Type uint8
		Data string
	}
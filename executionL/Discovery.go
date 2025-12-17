package main 

import (
	"net"
	"fmt"
	"bytes"
	"encoding/gob"
)

const (
	PING uint8 = 0x10
	PONG uint8 = 0x20
)
type Ping_Pong struct {
	From *Node
	To *Node
	Type uint8
}


func(n *Node) StartDiscovery() {
	udpAddr,_ := net.ResolveUDPAddr("udp" , n.Address)
 conn,err :=  net.ListenUDP("udp",udpAddr)
if err != nil {panic(err)} 
go func(){
	for {
		buf := make([]byte ,1024)
	num,addr,_ := 	conn.ReadFromUDP(buf)
	msg  := DecodeMsg(buf[:num])
	 switch msg.Type {
	 case PING:
		go n.HandlePING(msg,addr,conn)
	 case PONG:
		go n.HandlePONG(msg)
	 }
	 
	}
}()
}

func(n *Node) HandlePING(msg Ping_Pong ,addr *net.UDPAddr , conn *net.UDPConn) {
	if msg.To != n { fmt.Println("not the receiver node of ping")} 
go n.HandleKademliaIncoming(msg.From)
go n.SendPong(addr, msg.From , conn )
//can implement a dial also , if required, left for now 
}
func(n *Node) SendPong(addr *net.UDPAddr , pongReceiver *Node,conn *net.UDPConn) {
	pong := Ping_Pong{n,pongReceiver,PONG}
	_,err := conn.WriteToUDP(pong.SerializePingPong(),addr)
 if err != nil {panic(err)}
}
func(msg Ping_Pong) SerializePingPong() []byte {
var buf bytes.Buffer
encoder := gob.NewEncoder(&buf)
encoder.Encode(msg)
return buf.Bytes()
}
func(n *Node) HandlePONG(pong Ping_Pong) {
go n.HandleKademliaIncoming(pong.From)
//can iplement any other verification or check to establish a connection 
}
func DecodeMsg(payload []byte) Ping_Pong {
	buff := bytes.NewBuffer(payload)
	decoder := gob.NewDecoder(buff)
	var msg Ping_Pong
	decoder.Decode(&msg)
	return msg
}
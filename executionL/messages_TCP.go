package main 

 import ("encoding/gob"
"bytes"
"fmt"
"net")

const (
	MSG_HELLO uint8 = 0x01
	MSG_RET_HELLO uint8 = 0x02
)
//messaging using tcp , implementing a hello immediately after getting connected 
func(msg *Message ) Serialize() []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	encoder.Encode(msg)
	return buf.Bytes()
}
func (n *Node) SendATypeMsgToAPeer(msg Message , conn net.Conn ){
if n.Peers[conn.RemoteAddr().String()] == nil {
	panic("not a peer")
}
	sereializedMsg := msg.Serialize()

_,err :=  conn.Write(sereializedMsg)
if err!= nil {
	panic(err)
}
fmt.Println("msg.sent")
}

func(n *Node) HandleIncomingMsgsforPeers(conn net.Conn , adddata string) {
	fmt.Println(adddata)
if adddata == "msg handleing strtd via accept" {
	wg.Done()
}
	for {
		buf := make([]byte,2048)
		num,err := conn.Read(buf)
		
		if err != nil {panic(err)}
		fmt.Println(num)
		buff := bytes.NewBuffer(buf[:num])
		decoder := gob.NewDecoder(buff)
		var Msg struct{
			Type uint8
			Data string
		}
		decoder.Decode(&Msg)
		switch Msg.Type {
		case MSG_HELLO :
			fmt.Println(Msg.Data)
			go n.HandleHello(Msg , conn )
			fmt.Println("called handle hello..")
	
		case MSG_RET_HELLO:
			fmt.Println("msg sent by" , Msg.Data)
			go n.HandleReturnHello(Msg , conn)
		}

	}
}	
 var ch chan bool = make(chan bool)
func(n *Node) HandleHello(msg Message , conn net.Conn) {
	
	fmt.Println("handling hello...")
	
	pr := n.Peers[conn.RemoteAddr().String()]
	pr.Id = msg.Data
	n.IdToPeerMap[msg.Data] = &conn
	go n.ReturnHello(conn)
	
}
func(n *Node) ReturnHello(conn net.Conn) {
msg := Message{MSG_RET_HELLO , n.Id}
	n.SendATypeMsgToAPeer(msg, conn)
}
func(n *Node) HandleReturnHello(msg Message , conn net.Conn) {
	fmt.Println("handling return hello for " , n.Id)
 pr := n.Peers[conn.RemoteAddr().String()]
  pr.Id = msg.Data
  n.IdToPeerMap[msg.Data] = &conn
  ch <- true
}
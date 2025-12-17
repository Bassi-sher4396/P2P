package main 
import ("net"
"fmt"
"log"
"sync"
)

var wg sync.WaitGroup 

func (n *Node) CreateListener(network string)(<-chan bool , error){
peerFilled := make(chan bool )
	ln , err := net.Listen("tcp",n.Address)
if err != nil {
fmt.Println(err)
return peerFilled,err
}
n.Listener = ln 
fmt.Println("listening on " ,n.Listener.Addr())
	

go func() {
	
	fmt.Println("...")
for {
	fmt.Println("xxx")
	
	fmt.Println("lll")
	wg.Add(1)
	conn,err:=  n.Listener.Accept()
	fmt.Println("chchchchchhc")
	if err != nil {
		log.Fatal(err)
		continue
	} 
	if conn != nil {
		
	n.Peers[conn.RemoteAddr().String()] = &Peer{"",conn}

		go n.HandleIncomingMsgsforPeers(conn , "msg handleing strtd via accept")

	peerFilled <- true
	fmt.Println("NEw connection established")
}}
}()

return peerFilled,nil
}
func (n *Node) DialToOtherNode(network , address string) error {
conn,err:=	net.Dial(network , address )
if err != nil {fmt.Println(err)
return err}
fmt.Println("connected")
n.Peers[conn.RemoteAddr().String()] = &Peer{"",conn}
go n.HandleIncomingMsgsforPeers(conn,"msg handling strtd via dial")
msg := Message{	MSG_HELLO,n.Id} 
wg.Wait()
n.SendATypeMsgToAPeer(msg,conn)
return nil
}
package main 
import ("fmt"
"time"
"net"
)

func main() { 

	nodeA := Node{Id:"NodeA",Address:"127.0.0.1:3000",Peers :make(map[string]*Peer),IdToPeerMap :make(map[string]*net.Conn),Kademlia : make([]*Node,k_length),}
// if _,_,err :=	nodeA.CreateListener("tcp") ; err != nil {
// 	panic(err)
// }

//	fmt.Println(nodeA.Peers,nodeA.Listener.Addr())
	nodeB := Node{Id:"NodeB",Address:"127.0.0.1:3001",Peers:make(map[string]*Peer),IdToPeerMap : make(map[string]*net.Conn),Kademlia : make([]*Node,k_length),}
peerFilled,errr :=	nodeB.CreateListener("tcp") 
	
if errr != nil {
	panic(errr)
}
time.Sleep(2*time.Second)
// 	fmt.Println(nodeB.Peers,nodeB.Listener.Addr())
	
	
	fmt.Println("::")
if err :=	nodeA.DialToOtherNode("tcp","127.0.0.1:3001"); err!= nil {
	panic(err)
} else {
	fmt.Println("succesfully connected")
}
//time.Sleep(2*time.Second)
<-peerFilled
	fmt.Println(nodeA.Peers ,nodeB.Peers)
//time.Sleep(5*time.Second)
	// var mssg = Message{
	// 	MSG_HELLO,
	// 	"hello"	}
	//nodeA.SendATypeMsgToAPeer(mssg,nodeA.Peers)
	<-ch
	fmt.Println(nodeA.IdToPeerMap["NodeB"])
}
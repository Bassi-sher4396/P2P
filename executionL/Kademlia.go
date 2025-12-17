package main 

import "fmt"
const k_length = 16

//our kademlia is a bit different , it keeps nodes with high numbers of peers while discarding the nodes with a lower number 
func(n *Node) HandleKademliaIncoming(newentrant *Node) {
	n.Mu.Lock()
	defer n.Mu.Unlock()
	if len(n.Kademlia) < k_length {
		
		n.Kademlia = append(n.Kademlia, newentrant)
	}
		if len(n.Kademlia) == k_length {
		i,t :=	n.ChecktheLowestOne()
		if len(newentrant.Peers) > t {

n.Kademlia = append(n.Kademlia[:i],n.Kademlia[i+1:]...)
n.Kademlia = append(n.Kademlia , newentrant)

		}else{
			fmt.Println("kademlia full with better nodes")
		}
	}
}
func(n *Node) ChecktheLowestOne() ( int, int ) {
k :=  n.Kademlia 
var i int 
	var t int
	var lowestCountOfPeers int
	
	for j,node := range k {
   t = len(node.Peers)
   if t == 0 {
	 return j,0
   }
   // to initiate the variable 
   if lowestCountOfPeers == 0 { 
	lowestCountOfPeers = t
	i = j 
   }
   if t < lowestCountOfPeers {
	lowestCountOfPeers = t
	i=j 
   }
	}
	return i,lowestCountOfPeers
}
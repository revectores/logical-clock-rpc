package main

import (
	"log"
	"net"
	"net/rpc"
)

type NodeService struct{}

var vector_clock []int
var cur_index int

type InitTuple struct {
	NodeCount int
	NodeIndex int
}

func max(int a, int b) int {
	if a > b {
		return a
	}
	return b
}

func vmax(int vec1[], int vec2[]) int[] {

	if vec1.len() {
  		panic(err)
	}

	for i := range vec1.len()

}



func (p *NodeService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func (p *NodeService) InitVector(initTuple InitTuple, reply *int) error {
	vector_clock, err = make([]int, InitTuple.NodeCount)
	cur_index = InitTuple.NodeIndex

	*reply = 0
	return nil
}

func (p *NodeService) UpdateState(_ int, reply *int) error {
	vector_clock[cur_index]++

	*reply = 0
	return nil
}

func (p *NodeService) Message(vector_msg int[], reply *int) error {


	*reply = 0
	return nil
}

func (p *NodeService) SendMessage(_ int, reply *int) error {
	vector_clock[cur_index]++

	*reply = 0
	return nil
}



func main() {
	rpc.RegisterName("NodeService", new(NodeService))

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

    for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

        go rpc.ServeConn(conn)
    } 
}

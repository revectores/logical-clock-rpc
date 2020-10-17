
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"strconv"
)

type Addr struct {
	IP [4]uint8
	Port uint16
}

var Id2addr map[int]Addr


var port uint16
var history []string
var thisId int
var initialized bool
var nodeMode int

var LOCALHOST = [4]uint8{127, 0, 0, 1}
const DEV_MODE = 0
const PROD_MODE = 1

func getAddrString(id int) string {
	addr := Id2addr[id]
	addrString := fmt.Sprintf("%d.%d.%d.%d:%d", addr.IP[0], addr.IP[1],
		addr.IP[2], addr.IP[3], addr.Port)
	return addrString
}

func rpcDial(id int) *rpc.Client {
	addr := getAddrString(id)
	client, err := rpc.Dial("tcp", addr)

	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}


type NodeService struct{}

func (p *NodeService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func (p *NodeService) Init(id2addr map[int]Addr, reply *int) error {
	Id2addr = id2addr
	nodeMode = DEV_MODE
	for _, addr := range id2addr {
		if addr.IP != LOCALHOST {
			nodeMode = PROD_MODE
			break
		}
	}

	switch nodeMode {
	case DEV_MODE:
		thisId = -1
		for id, addr := range id2addr {
			if addr.Port == port {
				thisId = id
				break
			}
		}
	case PROD_MODE:
		// TOOD: PROD_MODE
	}

	*reply = thisId
	return nil
}

func (p *NodeService) GetHistory(_ int, reply *[]string) error {

	*reply = history
	return nil
}

func (p *NodeService) UpdateState(_ int, reply *int) error {
	var number = 0
	for i:=0; i<len(history);i++{
		if history[0] == string(thisId){
			number++
		}
	}
	var name = string(thisId)+"."+string(number)
	history = append(history, name)
	*reply = 0
	return nil
}

func (p *NodeService) Message(history []string, reply *int) error {
	var number = 0
	for i:=0; i<len(history);i++{
		if history[0] == string(thisId){
			number++
		}
	}
	var name = string(thisId)+"."+string(number)
	history = append(history, name)
	*reply = 0
	return nil
}

func (p *NodeService) SendMessage(id int, reply *int) error {
	client := rpcDial(id)

	var i int
	err := client.Call("NodeService.Message", history, &i)

	if err != nil {
		log.Fatal(err)
	}

	*reply = 0
	return nil
}



func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: ./node [port]")
		return;
	}

	rpc.RegisterName("NodeService", new(NodeService))

	port_, _ := strconv.Atoi(args[1])
	port = uint16(port_)

	listener, err := net.Listen("tcp", ":" + args[1])

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

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

var Object []string
var version map[string][]int
var port uint16
var clockVector []int
var thisId int
var initialized bool
var nodeMode int

var LOCALHOST = [4]uint8{127, 0, 0, 1}
const DEV_MODE = 0
const PROD_MODE = 1


func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func vmax(vec1 []int, vec2 []int) []int {
	if len(vec1) != len(vec2) {
		panic("")
	}

	updatedVector := make([]int, len(vec1))

	for i := 0; i < len(vec1); i++ {
		if vec1[i] > vec2[i] {
			updatedVector[i] = vec1[i]
		} else {
			updatedVector[i] = vec2[i]
		}
	}

	return updatedVector
}

func versionmax(ver1 map[string][]int,ver2 map[string][]int)map[string][]int{
	for _, n := range Object{
		ver1[n] = vmax(ver1[n],ver2[n])
	}
	return ver1
}

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

type Obj struct {
	Id2addr map[int]Addr
	Object  []string
}
func (p *NodeService) Init(init Obj, reply *int) error {
	Id2addr = init.Id2addr
	Object = init.Object
	version =make(map[string][]int,len(Object))
	for _, n := range Object {
		version[n] = make([]int, len(Id2addr))
	}

	nodeMode = DEV_MODE
	for _, addr := range Id2addr {
		if addr.IP != LOCALHOST {
			nodeMode = PROD_MODE
			break
		}
	}

	switch nodeMode {
	case DEV_MODE:
		thisId = -1
		for id, addr := range Id2addr {
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

func (p *NodeService) GetVector(_ int, reply *map[string][]int) error {
	*reply = version
	return nil
}

func (p *NodeService) UpdateState(str string, reply *int) error {
	for _, n := range Object {
		if str != n {
			*reply = -1
		} else {
			version[str][thisId]++
			*reply = 0
			break
		}
	}
	return nil
}

func (p *NodeService) Message(versionMsg map[string][]int, reply *int) error {
	for _, n := range Object{
		version[n][thisId]++
	}
	version=versionmax(version,versionMsg)
	*reply = 0
	return nil
}


func (p *NodeService) SendMessage(id int, reply *int) error {
	client := rpcDial(id)

	var i int;
	err := client.Call("NodeService.Message", version, &i)

	if err != nil {
		log.Fatal(err)
	}

	*reply = 0;
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

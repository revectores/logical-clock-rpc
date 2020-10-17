package controller

import (
	"os"
	"fmt"
	"log"
	"net/rpc"
	"encoding/json"
	"io/ioutil"
)


type Addr struct {
	IP [4]uint8
	Port uint16
}

var Id2addr map[int]Addr


func LoadConfigure() {
	id2addrJSON, _ := ioutil.ReadFile("../../conf/id2addr.json")
	json.Unmarshal(id2addrJSON, &Id2addr)
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


func Init(id int) int {
	client := rpcDial(id)

	var thisId int
	err := client.Call("NodeService.Init", Id2addr, &thisId)

	if err != nil {
		log.Fatal(err)
	}

	return thisId
}


func UpdateState(id int) int {
	client := rpcDial(id)

	var res int
	err := client.Call("NodeService.UpdateState", 0, &res)

	if err != nil {
		log.Fatal(err)
	}

	return res
}


func SendMessage(senderId int, receiverId int) int {
	client := rpcDial(senderId)

	var res int
	err := client.Call("NodeService.SendMessage", receiverId, &res)

	if err != nil {
		log.Fatal(err)
	}

	return res
}


func GetVector(id int) []int {
	client := rpcDial(id)

	var clockVector []int;
	err := client.Call("NodeService.GetVector", 0, &clockVector)

	if err != nil {
		log.Fatal(err)
	}

	return clockVector
}


func main() {
	client, err := rpc.Dial("tcp", "localhost:" + os.Args[1])
	fmt.Println(os.Args[1])
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// var reply string
	var i int
	err = client.Call("NodeService.UpdateState", 0, &i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(i)
}

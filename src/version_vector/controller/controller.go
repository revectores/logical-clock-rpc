package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
)


type Addr struct {
	IP [4]uint8
	Port uint16
}

var Id2addr map[int]Addr
type Obj struct {
	Id2addr map[int]Addr
	Object []string
}

var Object []string
var Character Obj
func LoadConfigure() Obj {
	id2addrJSON, _ := ioutil.ReadFile("../../conf/id2addr.json")
	objectJSON, _ := ioutil.ReadFile("../../conf/object.json")
	json.Unmarshal(id2addrJSON, &Id2addr)
	json.Unmarshal(objectJSON,&Object)
	Character.Id2addr = Id2addr
	Character.Object = Object
	return Character
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


func Init(id int)int{
	client := rpcDial(id)

	var i int
	err := client.Call("NodeService.Init", Character, &i)

	if err != nil {
		log.Fatal(err)
	}

	return i
}


func UpdateState(id int, str string)int{
	client := rpcDial(id)

	var i int
	err := client.Call("NodeService.UpdateState", str, &i)

	if err != nil {
		log.Fatal(err)
	}
	return i
}


func SendMessage(senderId int, receiverId int) int{
	client := rpcDial(senderId)

	var i int
	err := client.Call("NodeService.SendMessage", receiverId, &i)

	if err != nil {
		log.Fatal(err)
	}

	return i
}


func GetVector(id int) map[string][]int{
	client := rpcDial(id)
	var version map[string][]int
	err := client.Call("NodeService.GetVector", 0, &version)

	if err != nil {
		log.Fatal(err)
	}

	return version
}


func main() {
	client, err := rpc.Dial("tcp", "localhost:" + os.Args[1])
	fmt.Println(os.Args[1])
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// var reply string
	var i int
	str := "x"
	err = client.Call("NodeService.UpdateState", str, &i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(i)
}

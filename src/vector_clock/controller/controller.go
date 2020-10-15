package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
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

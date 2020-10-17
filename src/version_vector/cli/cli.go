package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
	"strconv"
	"../controller"
)


func createNodes(id2addr map[int]controller.Addr){
	fmt.Println(len(id2addr))
	
	for _, addr := range id2addr {
		cmd := exec.Command("../node/node", strconv.Itoa(int(addr.Port)))
		go cmd.Run()
		fmt.Println(cmd)
	}

	/*
	for id := range id2addr {
		controller.Init(id)
	}*/
}


func validate(params []string) {
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	controller.LoadConfigure()
	fmt.Println(controller.Id2addr)
	fmt.Println(controller.Object)
	for {
		fmt.Print(">> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		params := strings.Fields(cmd)

		switch params[0] {
		case "c", "create":
			createNodes(controller.Id2addr)
		case "i", "init":
			if len(params) != 2 {
				fmt.Println("usage: i(nit) [node-id]")
				continue
			}
			id, _ := strconv.Atoi(params[1])
			fmt.Println(controller.Init(id))
		case "g", "get":
			if len(params) != 2 {
				fmt.Println("usage: g(et) [node-id]")
				continue
			}
			id, _ := strconv.Atoi(params[1])
			fmt.Println(controller.GetVector(id))
		case "u", "update":
			if len(params) != 3 {
				fmt.Println("usage: u(pdate) [node-id] [object-name]")
				continue
			}
			id, _ := strconv.Atoi(params[1])
			fmt.Println(controller.UpdateState(id,params[2]))
		case "s", "send":
			if len(params) != 3 {
				fmt.Println("usage: s(end) [sender-node-id] [receiver-node-id]")
			}
			senderId, _ := strconv.Atoi(params[1])
			receiverId, _ := strconv.Atoi(params[2])
			fmt.Println(controller.SendMessage(senderId, receiverId))
		case "e", "exit":
			return;
		default:
		}
	}
}

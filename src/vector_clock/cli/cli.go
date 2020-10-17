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


func createNodes(id2addr map[int]controller.Addr) []*exec.Cmd {
	var nodes []*exec.Cmd
	
	for _, addr := range id2addr {
		node := exec.Command("../node/node", strconv.Itoa(int(addr.Port)))
		go node.Run()

		nodes = append(nodes, node)
		fmt.Println(node)
	}

	return nodes;
	/*
	for id := range id2addr {
		controller.Init(id)
	}*/
}


func validate(params []string) {
}


func main() {
	var nodes []*exec.Cmd

	reader := bufio.NewReader(os.Stdin)
	controller.LoadConfigure()
	for {
		fmt.Print(">> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		params := strings.Fields(cmd)

		if len(params) == 0 {
			continue
		}

		switch params[0] {
		case "c", "create":
			nodes = createNodes(controller.Id2addr)
		case "i", "init":
			if len(params) != 2 {
				fmt.Println("usage: i(nit) [node-id]")
				continue
			}
			id, _ := strconv.Atoi(params[1]) 
			thisId := controller.Init(id)
			fmt.Println(thisId)
		case "g", "get":
			if len(params) != 2 {
				fmt.Println("usage: g(et) [node-id]")
				continue
			}
			id, _ := strconv.Atoi(params[1])
			clockVector := controller.GetVector(id)
			fmt.Println(clockVector)
		case "u", "update":
			if len(params) != 2 {
				fmt.Println("usage: u(pdate) [node-id]")
				continue
			}
			id, _ := strconv.Atoi(params[1])
			res := controller.UpdateState(id)
			fmt.Println(res)
		case "s", "send":
			if len(params) != 3 {
				fmt.Println("usage: s(end) [sender-node-id] [receiver-node-id]")
			}
			senderId, _ := strconv.Atoi(params[1])
			receiverId, _ := strconv.Atoi(params[2])
			res := controller.SendMessage(senderId, receiverId)
			fmt.Println(res)
		case "e", "exit":
			for _, node := range nodes {
				node.Process.Kill()
			}
			return;
		default:
		}
	}
}

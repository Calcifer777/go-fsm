package main

import (
	"fmt"
	"log"

	"github.com/calcifer777/go-fsm/fsm"
)

func main() {
	node := fsm.NewPassNode("start", "step1", false)
	// fmt.Println(node)
	fmt.Println("ID: ", node.ID())
	node.Run(map[string]string{})

	node2 := fsm.NewPassNode("step1", "end", false)
	node3 := fsm.NewPassNode("end", "", true)
	nodes := make([]fsm.Node, 0)
	nodes = append(nodes, node, node2, node3)
	fmt.Println(nodes)

	myFsm := fsm.NewFsm(nodes)
	fmt.Println(myFsm)

	runner := fsm.NewFsmRunner(*myFsm, "start")
	fmt.Println(runner)

	out, err := runner.Run(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("output context", out)
	fmt.Println(runner)

}

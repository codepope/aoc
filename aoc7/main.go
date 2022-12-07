package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

 type Node struct {
	size int
	name string
	children map[string]Node
	parent *Node
}

func (n *Node) addChild(name string, size int) Node {

	node:=&Node{
		size:     size,
		name:     name,
		children: map[string]Node{},
		parent:   &Node{},
	}
	node.

	n.children[name]=node
}

func main() {
	if len(os.Args) != 2 {
		panic("specify file")
	}

	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	root:=&Node{
		size:     -1,
		name:     "/",
		children: map[string]Node{},
		parent:   nil,
	}

	currnode:=root
	
	for scanner.Scan() {
		line := scanner.Text()

		if string.StartsWith(line,"$") {
			// Command. let's parse
			if line=="$ cd /" {
				currnode=root
			} else if 
		}
	}
}

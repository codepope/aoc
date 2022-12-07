package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	size     int
	name     string
	children map[string]*Node
	parent   *Node
}

func (n *Node) addChild(name string, size int) {

	node := &Node{
		size:     size,
		name:     name,
		children: make(map[string]*Node),
		parent:   n,
	}

	n.children[name] = node
}

func (n *Node) calcSize() {
	if n.size != -1 {
		return
	}

	size := 0
	for _, child := range n.children {
		child.calcSize()
		size += child.size
	}

	n.size = size
}

func (n *Node) targetSize(free int, target int) []*Node {
	if len(n.children) == 0 {
		return nil
	}

	targets := make([]*Node, 0)

	for _, child := range n.children {
		r := child.targetSize(free, target)
		if r != nil {
			targets = append(targets, r...)
		}
	}

	if (free + n.size) > target {
		targets = append(targets, n)
	}

	return targets
}

type BySize []*Node // for sorting

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

	root := &Node{
		size:     -1,
		name:     "/",
		children: map[string]*Node{},
		parent:   nil,
	}

	currnode := root

	for scanner.Scan() {
		line := scanner.Text()

		cmd := strings.Split(line, " ")

	cmdprocess:
		if cmd[0] == "$" && cmd[1] == "cd" {
			if cmd[2] == "/" {
				currnode = root
			} else if cmd[2] == ".." {
				currnode = currnode.parent
			} else {
				currnode = currnode.children[cmd[2]]
			}
		} else if cmd[0] == "$" && cmd[1] == "ls" {
			for scanner.Scan() {
				line = scanner.Text()

				cmd = strings.Split(line, " ")
				if cmd[0] == "$" {
					goto cmdprocess
				}
				if cmd[0] == "dir" {
					currnode.addChild(cmd[1], -1)
				} else {
					size, _ := strconv.Atoi(cmd[0])
					currnode.addChild(cmd[1], size)
				}
			}
		}
	}

	root.calcSize()

	targets := root.targetSize(70000000-root.size, 30000000)

	sort.Slice(targets, func(i, j int) bool {
		return targets[i].size < targets[j].size
	})

	// for _, target := range targets {
	// 	fmt.Printf("%10s %7d\n", target.name, target.size)
	// }

	fmt.Printf("%10s %7d\n", targets[0].name, targets[0].size)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cycle int
var regX int

func tick(value int) {
	modcycle := cycle % 40

	//fmt.Println("In cycle ", cycle, " X is ", regX, "modcycle=", modcycle)

	if modcycle == regX || modcycle == regX+1 || modcycle == regX+2 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if modcycle == 0 {
		fmt.Println("")
	}
	if value == 0 {
		//fmt.Println("End cycle ", cycle)
		cycle++
		return
	}
	regX += value
	//fmt.Println("End cycle ", cycle, "Adding ", value, " X is now ", regX)
	cycle++

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

	cycle = 1
	regX = 1

	for scanner.Scan() {

		parts := strings.Split(scanner.Text(), " ")
		if len(parts) == 2 {
			tick(0)
			value, _ := strconv.Atoi(parts[1])
			tick(value)
		} else {
			tick(0)
		}
	}

}

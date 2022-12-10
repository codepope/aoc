package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cycle int
var interesting int
var regX int
var valueCount = 6
var valueTotal = 0

func tick() {
	cycle++
	if cycle == interesting {
		fmt.Println("****", cycle, cycle*regX)
		valueCount--
		if valueCount >= 0 {
			valueTotal += cycle * regX
		}
		interesting += 40
	}
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

	cycle = 0
	interesting = 20
	regX = 1

	for scanner.Scan() {
		tick()

		parts := strings.Split(scanner.Text(), " ")
		if len(parts) == 2 {
			value, _ := strconv.Atoi(parts[1])
			tick()
			regX += value

		} else {
			//fmt.Println("noop")
		}
	}

	fmt.Println(valueTotal)
}

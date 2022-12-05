package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		pair1 := strings.Split(numbers[0], "-")
		np10, _ := strconv.Atoi(pair1[0])
		np11, _ := strconv.Atoi(pair1[1])
		pair2 := strings.Split(numbers[1], "-")
		np20, _ := strconv.Atoi(pair2[0])
		np21, _ := strconv.Atoi(pair2[1])
		if (np10 >= np20 && np11 <= np21) ||
			(np20 >= np10 && np21 <= np11) {
			fmt.Println("Enclosed")
			total += 1
		}
	}

	fmt.Println(total)
}

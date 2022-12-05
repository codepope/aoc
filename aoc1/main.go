package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.OpenFile("aoc1input.txt", os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	totals := make([]int, 0)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			totals = append(totals, total)
			total = 0
		} else {
			num, _ := strconv.Atoi(line)
			total += num
		}
	}

	sort.Ints(totals)

	totalthree := totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]

	fmt.Println(totalthree)
}

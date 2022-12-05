package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []byte

func (s stack) Push(v []byte) stack {
	return append(s, v...)
}

func (s stack) Pop(n int) (stack, []byte) {
	l := len(s)
	return s[:l-n], s[l-n : l]
}

func (s stack) Dump() {
	fmt.Printf("%d ", len(s))
	for n := 0; n < len(s); n++ {
		fmt.Printf("%c", s[n])
	}
	fmt.Println()
}

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	// Build the stacks
	// Last line is going to be the starting point
	lastline := len(lines) - 1
	laststring := lines[lastline]

	dock := make([]stack, 0)

	for i := 0; i < len(laststring); i++ {
		if laststring[i] != 32 {
			// We got a column....
			newstack := make(stack, 0)
			for j := lastline - 1; j >= 0; j-- {
				if len(lines[j]) > j {
					if (lines[j][i]) != 32 {
						newstack = newstack.Push([]byte{lines[j][i]})
					}
				}
			}

			dock = append(dock, newstack)
		}

	}
	for i, d := range dock {
		fmt.Print(i)
		fmt.Print(" ")
		d.Dump()
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		count, _ := strconv.Atoi(parts[1])
		fromcol, _ := strconv.Atoi(parts[3])
		tocol, _ := strconv.Atoi(parts[5])
		fmt.Printf("move %d from %d to %d\n", count, fromcol, tocol)

		tmp := []byte{}

		dock[fromcol-1], tmp = dock[fromcol-1].Pop(count)
		dock[tocol-1] = dock[tocol-1].Push(tmp)

		for i, d := range dock {
			fmt.Print(i)
			fmt.Print(" ")
			d.Dump()
		}

	}

}

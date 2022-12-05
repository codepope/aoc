package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findChars(s string, t string) string {
	for _, c := range t {
		if strings.ContainsRune(s, c) {
			return string(c)
		}
	}

	return ""
}

const charindex = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		splitpoint := len(line) / 2
		comp1 := line[:splitpoint]
		comp2 := line[splitpoint:]

		reschar := findChars(comp1, comp2)

		score := strings.Index(charindex, reschar) + 1

		total = total + score

	}

	fmt.Println(total)
}

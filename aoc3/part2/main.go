package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findChars(s string, t string, u string) string {
	for _, c := range s {
		if strings.ContainsRune(t, c) && strings.ContainsRune(u, c) {
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
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		fmt.Printf("%s %s %s\n", line1, line2, line3)

		common := findChars(line1, line2, line3)

		fmt.Println(common)

		score := strings.Index(charindex, common) + 1
		total += score

		// splitpoint := len(line) / 2
		// comp1 := line[:splitpoint]
		// comp2 := line[splitpoint:]

		// reschar := findChars(comp1, comp2)

		// score := strings.Index(charindex, reschar) + 1

		// total = total + score

	}

	fmt.Println(total)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func rps(index string, choice byte) int {
	if index[0] == choice {
		return 1 // rock
	}
	if index[1] == choice {
		return 2 // paper
	}
	if index[2] == choice {
		return 3 // scissors
	}
	return 0
}

func goal(them int, goal int) int {
	switch goal {
	case 1: // Lose
		switch them {
		case 1:
			return 3
		case 2:
			return 1
		case 3:
			return 2
		}
	case 2: // Draw
		return them
	case 3: // Win
		switch them {
		case 1:
			return 2
		case 2:
			return 3
		case 3:
			return 1

		}
	}
	panic("WTF")
}

func play(them int, mine int) int {

	if (mine == 1 && them == 3) ||
		(mine == 3 && them == 2) ||
		(mine == 2 && them == 1) {
		return mine + 6
	}
	if (them == 1 && mine == 3) ||
		(them == 3 && mine == 2) ||
		(them == 2 && mine == 1) {
		return mine + 0
	}

	return mine + 3
}

func main() {
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		them := rps("ABC", line[0])
		mine := rps("XYZ", line[2])

		myplay := goal(them, mine)

		fmt.Println(them, mine, play(them, mine))
		total = total + (play(them, myplay))
	}

	fmt.Println(total)
}

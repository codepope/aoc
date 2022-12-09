package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p Point) Move(dir string) Point {
	switch dir {
	case "U":
		return Point{p.X + 1, p.Y}
	case "D":
		return Point{p.X - 1, p.Y}
	case "L":
		return Point{p.X, p.Y - 1}
	case "R":
		return Point{p.X, p.Y + 1}
	}
	panic("unknown direction")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p Point) Follow(head Point) Point {
	dirX := p.X - head.X
	dirY := p.Y - head.Y

	fmt.Println(abs(dirX), abs(dirY))

	if (abs(dirX) == 1 && abs(dirY) == 0) || (abs(dirX) == 0 && abs(dirY) == 1) {
		return p
	}

	if abs(dirX) == 1 && abs(dirY) == 1 {
		return p
	}

	if abs(dirX) == 2 && abs(dirY) == 0 {
		return Point{p.X - dirX/2, p.Y}
	}

	if abs(dirY) == 2 && abs(dirX) == 0 {
		return Point{p.X, p.Y - dirY/2}
	}

	if abs(dirX) == 2 && abs(dirY) == 1 {
		return Point{p.X - dirX/2, p.Y - dirY}
	}

	if abs(dirY) == 2 && abs(dirX) == 1 {
		return Point{p.X - dirX, p.Y - dirY/2}
	}

	return p
}

func Dump(head Point, tail Point, start Point) {
	for n := 4; n >= 0; n-- {
		for m := 0; m <= 5; m++ {
			p := Point{n, m}
			if p == head {
				fmt.Print("H")
			} else if p == tail {
				fmt.Print("T")
			} else if p == start {
				fmt.Print("s")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
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

	head := Point{0, 0}
	tail := Point{0, 0}
	start := Point{0, 0}
	visits := make(map[Point]int)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		dir := parts[0]
		cnt, _ := strconv.Atoi(parts[1])

		fmt.Println(dir, cnt)
		for n := 1; n <= cnt; n++ {
			head = head.Move(dir)
			Dump(head, tail, start)
			tail = tail.Follow(head)
			Dump(head, tail, start)
			visits[tail]++
		}
	}

	fmt.Println(visits)
	fmt.Println(len(visits))
}

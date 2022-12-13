package main

// https://adventofcode.com/2022/day/12
// No credit to me on this, I seeked help and found a solution which made sense to me.
// Then I took it to bits and put it back together.
// Upside of this is I now get maps with structs as keys as a solution to sparse arrays.
// Oh and reading about path finding algorithms. - Funny how far you get with flatter algorithms
//

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type coord struct {
	x int
	y int
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	mymap := make([][]byte, 0)

	ps := make([]coord, 0)
	pe := coord{0, 0}

	j := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			row := make([]byte, len(line))
			for i := 0; i < len(line); i++ {
				if line[i] == 'S' || line[i] == 'a' {
					row[i] = 'a'
					ps = append(ps, coord{i, j})
				} else if line[i] == 'E' {
					row[i] = 'z'
					pe = coord{i, j}
				} else {
					row[i] = line[i]
				}
			}
			mymap = append(mymap, row)
			j++
		}
	}

	minSteps := math.MaxInt

	for i := 0; i < len(ps); i++ {
		steps := seek(mymap, ps[i], pe)
		if steps < minSteps {
			minSteps = steps
		}
	}

	fmt.Println("Min steps", minSteps)
}

func printMaps(mymap [][]byte) {
	for i := 0; i < len(mymap); i++ {
		for j := 0; j < len(mymap[0]); j++ {
			fmt.Print(string(mymap[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func getC(mymap [][]byte, p coord) byte {
	if p.y < 0 || p.x < 0 || p.y >= len(mymap) || p.x >= len(mymap[0]) {
		return 0
	}
	return mymap[p.y][p.x]
}

func seek(mymap [][]byte, ps coord, pe coord) int {
	seen := make(map[coord]bool)
	seen[ps] = true
	queue := []coord{ps}
	steps := 0
	var dirs = []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	found := false

loop:
	for len(queue) > 0 {
		k := len(queue)
		for i := 0; i < k; i++ {
			c := queue[0]
			queue = queue[1:]
			if c == pe {
				found = true
				break loop
			}
			for _, d := range dirs {
				nc := coord{c.x + d.x, c.y + d.y}
				if nc.x < 0 || nc.y < 0 || nc.x >= len(mymap[0]) || nc.y >= len(mymap) {
					continue
				}
				if seen[nc] {
					continue
				}
				isGreater := getC(mymap, nc) > getC(mymap, c)
				if isGreater && getC(mymap, nc)-getC(mymap, c) > 1 {
					continue
				}
				seen[nc] = true
				queue = append(queue, nc)
			}
		}
		steps++
	}

	if !found {
		return math.MaxInt
	}

	return steps
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	trees := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		treerow := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			treerow[i] = int(line[i]) - 48
		}
		trees = append(trees, treerow)
	}

	total := 0

	for i := 0; i <= len(trees)-1; i++ {
		for j := 0; j <= len(trees)-1; j++ {
			visible := scanTrees(i, j, trees)
			if visible {
				total++
			}
		}
	}

	fmt.Println(total)
}

func scanTrees(x, y int, trees [][]int) bool {
	if x == 0 || y == 0 || x == len(trees)-1 || y == len(trees)-1 {
		return true
	}

	nvisible := scanNorth(x, y, trees)
	evisible := scanEast(x, y, trees)
	wvisible := scanWest(x, y, trees)
	svisible := scanSouth(x, y, trees)

	return nvisible || evisible || wvisible || svisible
}

func scanNorth(x, y int, trees [][]int) bool {
	h := trees[x][y]
	for i := x - 1; i >= 0; i-- {
		if trees[i][y] >= h {
			return false
		}
	}
	return true
}

func scanSouth(x, y int, trees [][]int) bool {
	h := trees[x][y]
	for i := x + 1; i <= len(trees)-1; i++ {
		if trees[i][y] >= h {
			return false
		}
	}
	return true
}

func scanWest(x, y int, trees [][]int) bool {
	h := trees[x][y]
	for i := y - 1; i >= 0; i-- {
		if trees[x][i] >= h {
			return false
		}
	}
	return true
}

func scanEast(x, y int, trees [][]int) bool {
	h := trees[x][y]
	for i := y + 1; i <= len(trees)-1; i++ {
		if trees[x][i] >= h {
			return false
		}
	}
	return true
}

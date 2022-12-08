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

	max := 0

	for i := 0; i <= len(trees)-1; i++ {
		for j := 0; j <= len(trees)-1; j++ {
			score := scoreTrees(i, j, trees)
			if score > max {
				max = score
			}
		}
	}

	fmt.Println(max)
}

func scoreTrees(x, y int, trees [][]int) int {

	nvisible := scanNorth(x, y, trees)
	evisible := scanEast(x, y, trees)
	wvisible := scanWest(x, y, trees)
	svisible := scanSouth(x, y, trees)

	return nvisible * evisible * wvisible * svisible
}

func scanNorth(x, y int, trees [][]int) int {
	h := trees[x][y]
	score := 0
	for i := x - 1; i >= 0; i-- {
		score++
		if trees[i][y] >= h {
			return score
		}
	}
	return score
}

func scanSouth(x, y int, trees [][]int) int {
	h := trees[x][y]
	score := 0
	for i := x + 1; i <= len(trees)-1; i++ {
		score++
		if trees[i][y] >= h {
			return score
		}
	}
	return score
}

func scanWest(x, y int, trees [][]int) int {
	h := trees[x][y]
	score := 0
	for i := y - 1; i >= 0; i-- {
		score++
		if trees[x][i] >= h {
			return score
		}
	}
	return score
}

func scanEast(x, y int, trees [][]int) int {
	h := trees[x][y]
	score := 0
	for i := y + 1; i <= len(trees)-1; i++ {
		score++
		if trees[x][i] >= h {
			return score
		}
	}
	return score
}

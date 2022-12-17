package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run aoc14.go <input file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	lx := -1
	ux := -1
	ly := -1
	uy := -1

	xwidth := 1000
	yheight := 1000

	rocks := make([][]int, yheight)

	for i := range rocks {
		rocks[i] = make([]int, xwidth)
		for j := range rocks {
			rocks[i][j] = 0
		}
	}

	re := regexp.MustCompile(" -> ")
	for _, l := range lines {
		split := re.Split(l, -1)
		fmt.Println(split)
		px := -1
		py := -1
		for _, n := range split {
			var x int
			var y int
			fmt.Sscanf(n, "%d,%d", &x, &y)
			if x < lx || lx == -1 {
				lx = x
			}
			if x > ux || ux == -1 {
				ux = x
			}
			if y < ly || ly == -1 {
				ly = y
			}
			if y > uy || uy == -1 {
				uy = y
			}
			if px == x {
				// Horizontal
				sy := y
				ey := py
				if sy > ey {
					ty := ey
					ey = sy
					sy = ty
				}
				for n := sy; n <= ey; n++ {
					rocks[n][x] = 1
				}

			} else if py == y {
				// Horizontal
				sx := x
				ex := px
				if sx > px {
					tx := ex
					ex = sx
					sx = tx
				}
				for n := sx; n <= ex; n++ {
					rocks[y][n] = 1
				}
			}
			px, py = x, y
		}
	}

	for n := 0; n < xwidth; n++ {
		rocks[uy+2][n] = 1
	}

	printRocks(uy, lx, ux, rocks)

	done := false
	alldone := false
	sanddrop := 0
	for !alldone {
		spx := 500
		spy := 0
		if rocks[spy][spx] == 0 {
			done = false
			sanddrop++
			for !done {
				npx := spx
				npy := spy + 1

				if rocks[npy][npx] != 0 {
					if rocks[npy][npx-1] != 0 {
						if rocks[npy][npx+1] != 0 {
							rocks[spy][spx] = 2
							done = true
						} else {
							npx = npx + 1
						}
					} else {
						npx = npx - 1
					}
				}
				spx = npx
				spy = npy
				// if spy > uy {
				// 	alldone = true
				// 	done = true
				// }
			}
		} else {
			alldone = true
		}

	}
	printRocks(uy, lx, ux, rocks)

	fmt.Println(sanddrop)
}

func printRocks(uy int, lx int, ux int, rocks [][]int) {
	for i := 0; i <= uy+5; i++ {
		for j := lx - 2; j <= ux+2; j++ {
			switch rocks[i][j] {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

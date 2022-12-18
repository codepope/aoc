package main

// No credit to me on this solution, this is a reimplementation of another solution which
// uses some clumsy structs to work out whats where. It also made some terrible assumptions originally
// about how Manhattan Distances worked and, well, if you want to see the elegant inspiration for this
// it's at https://github.com/mnml/aoc/blob/main/2022/15/1.go
//
// On the upside, I learnt that appending copies and that continue to a label is super powerful and in
// no way whatsoever a goto on steroids. More like meth. :)
//

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Sensor struct {
	sx int
	sy int
	bx int
	by int
	md int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	if len(os.Args) > 3 {
		log.Fatal("need a filename argument")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("problem opening file ", err)
	}
	ln, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("need a line number")
	}
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)

	line := map[int]struct{}{}
	sensors := make([]Sensor, 0)
	for scan.Scan() {
		l := scan.Text()
		s := Sensor{}
		_, err := fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.sx, &s.sy, &s.bx, &s.by)
		if err != nil {
			log.Fatal("Ouchy ", err)
		}
		s.md = abs(s.sx-s.bx) + abs(s.sy-s.by)
		sensors = append(sensors, s)
		rd := s.md - abs(ln-s.sy)
		for x := s.sx - rd; x <= s.sx+rd; x++ {
			if !(s.bx == x && s.by == ln) {
				line[x] = struct{}{}
			}
		}
	}

	//fmt.Println(len(line))

	for y := 0; y <= ln*2; y++ {
	loop:
		for x := 0; x <= ln*2; x++ {
			for _, s := range sensors {
				dx := s.sx - x
				dy := s.sy - y
				if abs(dx)+abs(dy) <= s.md {
					x += s.md - abs(dy) + dx
					continue loop
				}
			}
			fmt.Println(x*4000000 + y)
			return
		}
	}
}

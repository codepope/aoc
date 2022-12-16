package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// type Packet struct {
// 	value int
// 	packets []Packet
// 	packetype string // value, array
// }

// func(p *Packet) Push(np Packet) {
// 	p.packets = append(p.packets, np)
// 	// fmt.Println("Pushed " + np.Value() + " onto " + p.Value())
// }

// func (p *Packet) Pop() Packet {
// 	t := p.packets[len(p.packets)-1]
// 	p.packets = p.packets[:len(p.packets)-1]
// 	//	fmt.Println("Popped " + t.Value() + " leaving " + p.Value())
// 	return t
// }

// func parse(r *bufio.Reader) Packet {
// 	var cp Packet
// 	done := false

//		for !done {
//			ch, _, err := r.ReadRune()
//			if err != nil {
//				if err == io.EOF {
//					done = true
//					continue
//				}
//				log.Fatal(err)
//			}
//			if ch == '[' {
//				p = parse(r)
//				cp.Push(p)
//			} else if unicode.IsDigit(ch) {
//				ch2, _, err := r.ReadRune()
//				if err != nil {
//					log.Fatal(err)
//				}
//				if unicode.IsDigit(ch2) {
//					p = &IntPacket{(int(ch)-'0')*10 + (int(ch2) - '0')}
//				} else {
//					p = &IntPacket{int(ch) - '0'}
//					r.UnreadRune()
//				}
//				pp.Push(p)
//			} else if ch == ']' {
//				done = true
//			} else {
//				//Drop char
//			}
//		}
//		return &pp
//	}
//
// //
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run aoc13.go <input file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	partcnt := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var l any
		json.Unmarshal([]byte(line), &l)

		scanner.Scan()
		line = scanner.Text()
		if line == "" {
			continue
		}
		var r any
		json.Unmarshal([]byte(line), &r)

		result := compare(l, r)
		partcnt++
		fmt.Println(partcnt, result, l, r)
		if result < 0 {
			total += partcnt
		}
	}

	fmt.Println(total)
}

func compare(l any, r any) int {

	la, lok := l.([]any)
	ra, rok := r.([]any)

	switch {
	case !lok && !rok:
		return int(l.(float64) - r.(float64))
	case !lok:
		la = []any{l}
	case !rok:
		ra = []any{r}
	}

	for i := 0; i < len(la) && i < len(ra); i++ {
		c := compare(la[i], ra[i])
		if c != 0 {
			return c
		}
	}

	return len(la) - len(ra)
}

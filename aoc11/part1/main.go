package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	monkeynumber int
	items        []int
	optype       string
	opold        bool
	opval        int
	testval      int
	truemonkey   int
	falsemonkey  int
	inspectcount int
}

func newMonkeyFromScanner(scan *bufio.Scanner) *Monkey {
	monkey := Monkey{}
	monkey.inspectcount = 0

	if !scan.Scan() {
		return nil
	} // Read Monkey line
	line := scan.Text()

	if line == "" { // discard blank lines
		if !scan.Scan() {
			return nil
		}
		line = scan.Text()
	}

	_, err := fmt.Sscanf(line, "Monkey %d:", &monkey.monkeynumber)
	if err != nil {
		log.Fatalf("format error for %s", line)
	}
	scan.Scan()
	line = scan.Text()
	parts := strings.Split(line, " ")
	for i := 4; i < len(parts); i++ {
		intval, _ := strconv.Atoi(strings.ReplaceAll(parts[i], ",", ""))
		monkey.items = append(monkey.items, intval)
	}
	scan.Scan()
	line = scan.Text()
	parts = strings.Split(line, " ")
	monkey.optype = parts[6]
	if parts[7] == "old" {
		monkey.opold = true
	} else {
		monkey.opold = false
		monkey.opval, _ = strconv.Atoi(parts[7])
	}

	scan.Scan()
	line = scan.Text()
	parts = strings.Split(line, " ")
	monkey.testval, _ = strconv.Atoi(parts[5])

	scan.Scan()
	line = scan.Text()
	parts = strings.Split(line, " ")

	monkey.truemonkey, _ = strconv.Atoi(parts[9])

	scan.Scan()
	line = scan.Text()
	parts = strings.Split(line, " ")
	monkey.falsemonkey, _ = strconv.Atoi(parts[9])

	return &monkey
}

func (m *Monkey) turn(monkeys []*Monkey) {

	for _, worryitem := range m.items {
		var worrylevel int
		var worryopval int
		m.inspectcount++
		if m.opold {
			worryopval = worryitem
		} else {
			worryopval = m.opval
		}
		if m.optype == "*" {
			worrylevel = worryitem * worryopval
		} else if m.optype == "+" {
			worrylevel = worryitem + worryopval
		} else {
			log.Fatalf("missing operand for money %d", m.monkeynumber)
		}
		worrylevel = worrylevel / 3

		if worrylevel%m.testval == 0 {
			monkeys[m.truemonkey].items = append(monkeys[m.truemonkey].items, worrylevel)
		} else {
			monkeys[m.falsemonkey].items = append(monkeys[m.falsemonkey].items, worrylevel)
		}
	}
	m.items = make([]int, 0)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("need filename to process")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("error opening file: %e", err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	monkeys := make([]*Monkey, 0)

	monkey := newMonkeyFromScanner(scanner)

	for monkey != nil {
		monkeys = append(monkeys, monkey)
		monkey = newMonkeyFromScanner(scanner)
	}

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.turn(monkeys)
		}

		for i := range monkeys {
			fmt.Println(monkeys[i])
		}
	}

	for i := range monkeys {
		fmt.Println(i, monkeys[i].inspectcount)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectcount > monkeys[j].inspectcount
	})

	for i := range monkeys {
		fmt.Println(i, monkeys[i].inspectcount)
	}
	fmt.Println(monkeys[0].inspectcount * monkeys[1].inspectcount)
}

//
// This is the <nth> pass at this AoC day and it was horrible. I ended up
// creating code (with a node structure and various other unmanagable nonsense)
// that was being a pig to debug and likely would never have worked well enough.
//
// And in frustration, I ended up reading this solution -
// https://www.reddit.com/r/adventofcode/comments/zkmyh4/comment/j03cwuj by
// Redditor Pyr0byt3 and the scales dropped from my eyes. Using json Unmarshal
// to act as a parser, directly using the any typed structures that it produced,
// using type assertions rather than switch statements (when there's not many
// types), and making the comparison function just return the final determining
// difference when it works out what the order was. Oh and that json
// unmarshalled numbers are not ints, they are float64s.
//
// A lot of learning on the back of this, around composing structs around
// interfaces also happened so, despite its infuriating nature, this has been
// worthwhile. Most of the credit goes to pyr0byt3; I read and rewrote as
// blindly as possible without indulging in giving myself a head injury as a
// reset. There's a whole load of debug left live, a different file reading
// process and ... well dammit it works so, onwards and upwards.
//
// Top tip: If your code is feeling too unweildy and unmanagable, it probably is
// and you may well have missed a simpler path. Always be prepared to reset.
//

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

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

	sortable := make([]any, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var l any
		json.Unmarshal([]byte(line), &l)
		sortable = append(sortable, l)

		scanner.Scan()
		line = scanner.Text()
		if line == "" {
			continue
		}
		var r any
		json.Unmarshal([]byte(line), &r)
		sortable = append(sortable, r)
	}

	div1 := "[[2]]"
	div2 := "[[6]]"
	var g1, g2 any
	json.Unmarshal([]byte(div1), &g1)
	json.Unmarshal([]byte(div2), &g2)
	sortable = append(sortable, g1, g2)

	sort.Slice(sortable, func(i, j int) bool { return compare(sortable[i], sortable[j]) < 0 })

	total := 1

	for i := 0; i < len(sortable); i++ {
		fmt.Println(i, fmt.Sprint(sortable[i]))
		if fmt.Sprint(sortable[i]) == "[[2]]" || fmt.Sprint(sortable[i]) == "[[6]]" {
			total *= (i + 1)
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

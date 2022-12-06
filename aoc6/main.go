package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type stack []byte

func main() {
	if len(os.Args) != 3 {
		panic("specify file and count")
	}

	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	cnt, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanRunes)

	//cnt := 14
	buf := make([]rune, cnt)
	ptr := cnt - 1
	index := 0

	for scanner.Scan() {
		scannedrune := rune(scanner.Text()[0])
		index += 1
		for i := 0; i < (cnt - 1); i++ {
			buf[i] = buf[i+1]
		}
		buf[cnt-1] = scannedrune

		if ptr >= 0 {
			ptr -= 1
		} else {
			fmt.Println(string(buf))
			match := false
			for j := 0; j <= (cnt - 1); j++ {
				for k := 0; k <= (cnt - 1); k++ {
					if j != k {

						if buf[j] == buf[k] {
							match = true
						}
					}
				}
			}
			if !match {
				fmt.Println(index)
				break
			}
		}
	}
}

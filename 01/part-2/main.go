package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	bytes, _ := ioutil.ReadFile("input.txt")

	var lines [951]int
	existing := make(map[int]bool, 100000)

	held := 0
	for i, v := range strings.Split(string(bytes), "\n") {
		toNumber, e := strconv.Atoi(v)
		if e != nil {
			break
		}
		lines[i] = toNumber
	}

	for {
		for _, n := range lines {
			held += n

			if _, exists := existing[held]; exists {

				print(strconv.Itoa(held))
				return
			}
			existing[held] = true
		}
	}
}

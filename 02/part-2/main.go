package main

import (
	"fmt"
	io "io/ioutil"
	str "strings"
)

func main() {

	f, _ := io.ReadFile("input.txt")
	lines := str.Split(string(f), "\n")

	// All lines
	for _, line := range lines {
		if checkAgainstAllOthers(lines, line) {
			fmt.Println(line)
			return
		}
	}
}

func checkAgainstAllOthers(all []string, current string) bool {
	for i := range all {
		if compareTwo(all[i], current) {
			return true
		}
	}
	return false
}

func compareTwo(str string, comparison string) bool {
	faults := 0
	for i := range str {
		if str[i] == comparison[i] {
			continue
		}
		faults++
	}

	if faults == 1 {
		return true
	}
	return false
}

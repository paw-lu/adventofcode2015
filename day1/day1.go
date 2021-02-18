package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("input.txt"))
}

func part1(input string) int {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	inputString := string(data)
	floor := 0

	for _, c := range inputString {
		if c == '(' {
			floor += 1
		} else if c == ')' {
			floor -= 1
		}
	}
	return floor
}

func part2(input string) int {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	inputString := string(data)
	floor := 0

	for i, c := range inputString {
		if c == '(' {
			floor += 1
		} else if c == ')' {
			floor -= 1
		}

		if floor == -1 {
			// indexes start at 1
			return i + 1
		}
	}
	return 0
}

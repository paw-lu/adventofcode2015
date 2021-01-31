package main

import "fmt"
import "io/ioutil"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)

	input := string(data)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	floor := 0

	for _, c := range input {
		if c == '(' {
			floor += 1
		} else if c == ')' {
			floor -= 1
		}
	}
	return floor
}

func part2(input string) int {
	floor := 0

	for i, c := range input {
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

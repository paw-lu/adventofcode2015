package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(part1("input.txt"))
}

func part1(input string) int {
	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	legList := []Leg{}
	for scanner.Scan() {
		parsedLeg := parseLeg(scanner.Text())
		legList = append(legList, parsedLeg)
	}
}

// Leg describes a single leg between two cities
type Leg struct {
	start    string
	dest     string
	distance int
}

func parseLeg(rawString string) Leg {
	return Leg{"Seattle", "San Francisco", 100}
}

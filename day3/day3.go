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
	inputRunes := []rune(string(data))

	houseMap := make(map[Coord]int)
	houseMap[Coord{0, 0}] = 1

	visitHouses(&houseMap, inputRunes)

	return len(houseMap)
}

func part2(input string) int {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}
	inputRunes := []rune(string(data))

	houseMap := make(map[Coord]int)
	houseMap[Coord{0, 0}] = 1

	input1 := make([]rune, 0, len(inputRunes)/2)
	for i := 0; i < len(inputRunes); i += 2 {
		input1 = append(input1, inputRunes[i])
	}

	input2 := make([]rune, 0, len(inputRunes)/2)
	for i := 1; i < len(inputRunes); i += 2 {
		input2 = append(input2, inputRunes[i])
	}

	visitHouses(&houseMap, input1)
	visitHouses(&houseMap, input2)

	return len(houseMap)
}

type Coord struct {
	x int
	y int
}

func (c *Coord) move(m rune) {
	switch m {
	case '^':
		c.y += 1
	case '>':
		c.x += 1
	case '<':
		c.x -= 1
	case 'v':
		c.y -= 1
	}
}

func visitHouses(hmap *map[Coord]int, input []rune) {
	position := Coord{0, 0}
	for _, c := range input {
		position.move(c)

		_, ok := (*hmap)[position]
		if ok {
			(*hmap)[position] += 1
		} else {
			(*hmap)[position] = 1
		}
	}
}

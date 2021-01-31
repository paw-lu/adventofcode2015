package main

import "fmt"
import "io/ioutil"

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := []rune(string(data))
	fmt.Println(part1(input))
	fmt.Println(part2(input))
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

func part1(input []rune) int {
	houseMap := make(map[Coord]int)
	houseMap[Coord{0, 0}] = 1

	visitHouses(&houseMap, input)

	return len(houseMap)
}

func part2(input []rune) int {
	houseMap := make(map[Coord]int)
	houseMap[Coord{0, 0}] = 1

	input1 := make([]rune, 0, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		input1 = append(input1, input[i])
	}

	input2 := make([]rune, 0, len(input)/2)
	for i := 1; i < len(input); i += 2 {
		input2 = append(input2, input[i])
	}

	visitHouses(&houseMap, input1)
	visitHouses(&houseMap, input2)

	return len(houseMap)
}

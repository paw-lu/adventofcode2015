package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("input.txt"))
}

func part1(input string) int {
	instructions := parseInput(input)
	g := BooleanGrid{}

	for _, inst := range instructions {
		g.executeBooleanInstruction(inst)
	}
	cnt := g.countBoolean()

	return cnt
}

func part2(input string) int {
	instructions := parseInput(input)
	g := IntegerGrid{}

	for _, inst := range instructions {
		g.executeIntegerInstruction(inst)
	}
	cnt := g.countInteger()

	return cnt
}

func parseInput(input string) []Instruction {
	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	instructions := []Instruction{}

	for scanner.Scan() {
		line := scanner.Text()
		instruction := parseLine(line)
		instructions = append(instructions, instruction)
	}

	return instructions
}

func parseLine(line string) Instruction {
	split := strings.Split(line, " ")
	minRange := strings.Split(split[len(split)-3], ",")
	maxRange := strings.Split(split[len(split)-1], ",")

	x1, _ := strconv.Atoi(minRange[0])
	y1, _ := strconv.Atoi(minRange[1])
	x2, _ := strconv.Atoi(maxRange[0])
	y2, _ := strconv.Atoi(maxRange[1])

	if strings.HasPrefix(line, "turn on") {
		return Instruction{"on", x1, y1, x2, y2}
	} else if strings.HasPrefix(line, "turn off") {
		return Instruction{"off", x1, y1, x2, y2}
	} else {
		return Instruction{"toggle", x1, y1, x2, y2}
	}
}

type BooleanGrid struct {
	coords [1000][1000]bool
}

type IntegerGrid struct {
	coords [1000][1000]int
}

type Instruction struct {
	action string
	x1     int
	y1     int
	x2     int
	y2     int
}

func (g *BooleanGrid) executeBooleanInstruction(inst Instruction) {
	for x := inst.x1; x <= inst.x2; x++ {
		for y := inst.y1; y <= inst.y2; y++ {
			switch inst.action {
			case "on":
				(&g.coords)[x][y] = true
			case "off":
				(&g.coords)[x][y] = false
			case "toggle":
				(&g.coords)[x][y] = !(&g.coords)[x][y]

			}
		}
	}
}

func (g *IntegerGrid) executeIntegerInstruction(inst Instruction) {
	for x := inst.x1; x <= inst.x2; x++ {
		for y := inst.y1; y <= inst.y2; y++ {
			switch inst.action {
			case "on":
				(&g.coords)[x][y]++
			case "off":
				if (&g.coords)[x][y] > 0 {
					(&g.coords)[x][y]--
				}
			case "toggle":
				(&g.coords)[x][y] += 2

			}
		}
	}
}

func (g BooleanGrid) countBoolean() int {
	i := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if g.coords[x][y] {
				i++
			}
		}
	}
	return i
}

func (g IntegerGrid) countInteger() int {
	i := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			i += g.coords[x][y]
		}
	}
	return i
}

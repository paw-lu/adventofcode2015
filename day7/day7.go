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
	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	circuit := make(map[string]Gate)

	for scanner.Scan() {
		sink, source := parseConnection(scanner.Text())
		circuit[sink] = source
	}

	wireMap := make(map[string]int)

	return traceSink("a", circuit, wireMap)
}

func part2(input string) int {
	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	circuit := make(map[string]Gate)

	for scanner.Scan() {
		sink, source := parseConnection(scanner.Text())
		circuit[sink] = source
	}

	wireMap := make(map[string]int)
	wireAOverride := traceSink("a", circuit, wireMap)

	revisedWireMap := make(map[string]int)
	revisedWireMap["b"] = wireAOverride

	return traceSink("a", circuit, revisedWireMap)
}

type Gate struct {
	operator string
	source0  string
	source1  string // null for RSHIFT/LSHIFT/INPUT
	param    int    // null for AND/OR
}

func parseConnection(instruction string) (string, Gate) {
	rawSplit := strings.Split(instruction, "->")
	sink := strings.TrimSpace(rawSplit[1])

	rawSource := strings.TrimSpace(rawSplit[0])
	source := Gate{}

	if strings.Contains(rawSource, "AND") {
		splitSource := strings.Split(rawSource, " AND ")
		source = Gate{
			operator: "AND",
			source0:  splitSource[0],
			source1:  splitSource[1],
		}
	} else if strings.Contains(rawSource, "OR") {
		splitSource := strings.Split(rawSource, " OR ")
		source = Gate{
			operator: "OR",
			source0:  splitSource[0],
			source1:  splitSource[1],
		}
	} else if strings.Contains(rawSource, "RSHIFT") {
		splitSource := strings.Split(rawSource, " RSHIFT ")
		shiftAmt, _ := strconv.Atoi(splitSource[1])

		source = Gate{
			operator: "RSHIFT",
			source0:  splitSource[0],
			param:    shiftAmt,
		}
	} else if strings.Contains(rawSource, "LSHIFT") {
		splitSource := strings.Split(rawSource, " LSHIFT ")
		shiftAmt, _ := strconv.Atoi(splitSource[1])

		source = Gate{
			operator: "LSHIFT",
			source0:  splitSource[0],
			param:    shiftAmt,
		}
	} else if strings.Contains(rawSource, "NOT") {
		splitSource := strings.Split(rawSource, "NOT ")
		source = Gate{
			operator: "NOT",
			source0:  splitSource[1],
		}
	} else {
		inputAmt, err := strconv.Atoi(rawSource)

		if err == nil {
			source = Gate{
				operator: "INPUT",
				param:    inputAmt,
			}
		} else {
			source = Gate{
				operator: "CONNECT",
				source0:  rawSource,
			}
		}
	}

	return sink, source
}

func traceSink(sink string, circuit map[string]Gate, wireMap map[string]int) int {
	g := circuit[sink]

	v, ok := wireMap[sink]
	if ok {
		return v
	}

	var wireValue int

	switch g.operator {
	case "AND":
		if g.source0 == "1" {
			wireValue = 1 & traceSink(g.source1, circuit, wireMap)
		} else {
			wireValue = traceSink(g.source0, circuit, wireMap) & traceSink(g.source1, circuit, wireMap)
		}
	case "OR":
		wireValue = traceSink(g.source0, circuit, wireMap) | traceSink(g.source1, circuit, wireMap)
	case "LSHIFT":
		wireValue = traceSink(g.source0, circuit, wireMap) << g.param
	case "RSHIFT":
		wireValue = traceSink(g.source0, circuit, wireMap) >> g.param
	case "NOT":
		wireValue = ^traceSink(g.source0, circuit, wireMap)
	case "CONNECT":
		wireValue = traceSink(g.source0, circuit, wireMap)
	case "INPUT":
		wireValue = g.param
	}

	wireMap[sink] = wireValue
	return wireValue
}

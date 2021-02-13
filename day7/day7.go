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
}

type Gate struct {
	operator string
	source0  string
	source1  string // null for RSHIFT/LSHIFT/INPUT
	param    int    // null for AND/OR
}

func part1(input string) int {
	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	circuit := make(map[string]Gate)

	for scanner.Scan() {
		sink, source := parseConnection(scanner.Text())
		circuit[sink] = source
	}

	return traceSink("a", circuit)
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

func traceSink(sink string, circuit map[string]Gate) int {
	fmt.Println("tracing wire ", sink)

	g := circuit[sink]

	switch g.operator {
	case "AND":
		return traceSink(g.source0, circuit) & traceSink(g.source1, circuit)
	case "OR":
		return traceSink(g.source0, circuit) | traceSink(g.source1, circuit)
	case "LSHIFT":
		return traceSink(g.source0, circuit) << g.param
	case "RSHIFT":
		return traceSink(g.source0, circuit) >> g.param
	case "NOT":
		return ^traceSink(g.source0, circuit)
	case "CONNECT":
		return traceSink(g.source0, circuit)
	case "INPUT":
		return g.param
	}

	return 0
}

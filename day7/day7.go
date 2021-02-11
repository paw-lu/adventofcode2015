package main

import "bufio"
import "fmt"
import "os"

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

	circuit = make(map[string]Gate)

	for scanner.Scan() {
		sink, source := parseConnection(scanner.Text())
		circuit[sink] = source
	}
	result := traceCircuit(circuit)

	return result["a"]
}

func parseConnection(instruction string) (sink string, source Gate) {
	rawSplit := strings.Split(instruction, "->")
	sink := strings.TrimSpace(rawSplit[1])

	rawSource := strings.TrimSpace(rawSplit[0])

	if strings.Contains(rawSource, "AND") {
		splitSource := strings.Split(rawSource, " AND ")
		source := Gate{
			operator: "AND",
			source0:  splitSource[0],
			source1:  splitSource[1],
		}
	} else if strings.Contains(rawSource, "OR") {
		splitSource := strings.Split(rawSource, " OR ")
		source := Gate{
			operator: "OR",
			source0:  splitSource[0],
			source1:  splitSource[1],
		}
	} else if strings.Contains(rawSource, "RSHIFT") {
		splitSource := strings.Split(rawSource, " RSHIFT ")
		source := Gate{
			operator: "RSHIFT",
			source0:  splitSource[0],
			param:    splitSource[1],
		}
	} else if strings.Contains(rawSource, "LSHIFT") {
		splitSource := strings.Split(rawSource, " LSHIFT ")
		source := Gate{
			operator: "RSHIFT",
			source0:  splitSource[0],
			param:    splitSource[1],
		}
	} else if strings.Contains(rawSource, "NOT") {
		splitSource := strings.Split(rawSource, " NOT ")
		source := Gate{
			operator: "NOT",
			source0:  splitSource[1],
		}
	} else {
		source := Gate{
			operator: "INPUT",
			param:    rawSource,
		}
	}

}

func traceCircuit(circuit map[string]Gate) map[string]int {

}

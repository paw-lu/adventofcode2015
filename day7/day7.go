package main

import "bufio"
import "fmt"
import "os"

func main() {
	fmt.Println(part1("input.txt"))
}

type Gate struct {
	operator string
	source1  string
	source2  string // null for RSHIFT/LSHIFT/INPUT
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

}

func traceCircuit(circuit map[string]Gate) map[string]int {

}

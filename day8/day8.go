package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(part1("input.txt"))
}

func part1(input string) int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	characterCount := 0
	memoryCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		memoryLine, _ := strconv.Unquote(line)

		characterCount += len(line)
		memoryCount += len(memoryLine)
	}

	return characterCount - memoryCount
}

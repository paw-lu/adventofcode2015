package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	input := "iwrupvqb"

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	i := 0
	for {
		testString := input + strconv.Itoa(i)
		hexString := getMd5Hex(testString)

		if hexString[:5] == "00000" {
			break
		}

		i++
	}

	return i
}

func part2(input string) int {
	i := 0
	for {
		testString := input + strconv.Itoa(i)
		hexString := getMd5Hex(testString)

		if hexString[:6] == "000000" {
			break
		}

		i++
	}

	return i
}

func getMd5Hex(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	paperNeeded := 0
	for scanner.Scan() {
		dims := parseToDimensions(scanner.Text())
		paperNeeded += dims.getRequiredArea()
	}

	fmt.Println(paperNeeded)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func MinOfInt(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

type Dimensions struct {
	l int
	w int
	h int
}

func (d Dimensions) getRequiredArea() int {
	basicArea := 2*d.l*d.w + 2*d.w*d.h + 2*d.h*d.l
	extraArea := MinOfInt(d.l*d.w, d.w*d.h, d.h*d.l)

	return basicArea + extraArea
}

func parseToDimensions(dimstr string) Dimensions {
	dimlist := strings.Split(dimstr, "x")

	l, e := strconv.Atoi(dimlist[0])
	check(e)

	w, e := strconv.Atoi(dimlist[1])
	check(e)

	h, e := strconv.Atoi(dimlist[2])
	check(e)

	return Dimensions{l, w, h}
}

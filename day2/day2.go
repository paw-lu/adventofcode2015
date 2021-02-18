package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	paperNeeded := 0
	for scanner.Scan() {
		dims := parseToDimensions(scanner.Text())
		paperNeeded += dims.getRequiredArea()
	}

	return paperNeeded
}

func part2(input string) int {
	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	ribbonNeeded := 0
	for scanner.Scan() {
		dims := parseToDimensions(scanner.Text())
		ribbonNeeded += dims.getRequiredRibbon()
	}

	return ribbonNeeded
}

type Dimensions struct {
	l int
	w int
	h int
}

func (d Dimensions) getRequiredArea() int {
	basicArea := 2*d.l*d.w + 2*d.w*d.h + 2*d.h*d.l
	extraArea := d.getSmallestArea()

	return basicArea + extraArea
}

func (d Dimensions) getSmallestArea() int {
	ds := []int{d.l * d.w, d.w * d.h, d.h * d.l}
	sort.Ints(ds)

	return ds[0]
}

func (d Dimensions) getRequiredRibbon() int {
	basicRibbon := d.getSmallestPerimeter()
	extraRibbon := d.l * d.w * d.h

	return basicRibbon + extraRibbon
}

func (d Dimensions) getSmallestPerimeter() int {
	ds := []int{2*d.l + 2*d.w, 2*d.w + 2*d.h, 2*d.h + 2*d.l}
	sort.Ints(ds)

	return ds[0]
}

func parseToDimensions(dimstr string) Dimensions {
	dimlist := strings.Split(dimstr, "x")

	l, _ := strconv.Atoi(dimlist[0])
	w, _ := strconv.Atoi(dimlist[1])
	h, _ := strconv.Atoi(dimlist[2])

	return Dimensions{l, w, h}
}

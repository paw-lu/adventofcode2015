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
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	ribbonNeeded := 0
	for scanner.Scan() {
		dims := parseToDimensions(scanner.Text())
		ribbonNeeded += dims.getRequiredRibbon()
	}

	fmt.Println(ribbonNeeded)
}

type Dimensions struct {
	l int
	w int
	h int
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

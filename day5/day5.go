package main

import "fmt"
import "bufio"
import "os"

func main() {
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("input.txt"))
}

func part1(input string) int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		if isPartOneNice(scanner.Text()) {
			i++
		}
	}

	return i
}

func isPartOneNice(s string) bool {
	runes := []rune(s)
	result := containsThreeVowels(runes) && hasTwiceInRow(runes) && noDisallowedPairs(runes)

	return result
}

func containsThreeVowels(r []rune) bool {
	vowelCount := 0
	for i, _ := range r {
		switch r[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCount++
		}
	}

	return vowelCount >= 3
}

func hasTwiceInRow(r []rune) bool {
	var currentLetter rune
	for _, nextLetter := range r {
		if nextLetter == currentLetter {
			return true
		} else {
			currentLetter = nextLetter
		}
	}

	return false
}

func noDisallowedPairs(r []rune) bool {
	for i, _ := range r {
		if i < len(r)-1 {
			switch r[i] {
			case 'a':
				if r[i+1] == 'b' {
					return false
				}
			case 'c':
				if r[i+1] == 'd' {
					return false
				}
			case 'p':
				if r[i+1] == 'q' {
					return false
				}
			case 'x':
				if r[i+1] == 'y' {
					return false
				}
			}
		} else {
			return true
		}
	}
	return true
}

func part2(input string) int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		if isPartTwoNice(scanner.Text()) {
			i++
		}
	}

	return i
}

func isPartTwoNice(s string) bool {
	runes := []rune(s)
	result := containsRepeatPair(runes) && containsSkippingLetter(runes)

	return result
}

func containsRepeatPair(r []rune) bool {
	type RunePair struct {
		first  rune
		second rune
	}

	letterPairs := []RunePair{}

	for i, _ := range r {
		if i < len(r)-1 {
			currentPair := RunePair{r[i], r[i+1]}
			// check if current pair has appeared previously
			for i, previousPair := range letterPairs {
				// exclude most recently-added pair
				if (i < len(letterPairs)-1) && (currentPair == previousPair) {
					return true
				}
			}
			letterPairs = append(letterPairs, currentPair)
		} else {
			break
		}

	}
	return false
}

func containsSkippingLetter(r []rune) bool {
	currentLetter := r[0]

	for i, _ := range r {
		if i < len(r)-2 {
			if currentLetter == r[i+2] {
				return true
			} else {
				currentLetter = r[i+1]
			}
		} else {
			return false
		}
	}

	return false
}

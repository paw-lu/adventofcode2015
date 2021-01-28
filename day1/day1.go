package main

import "fmt"
import "io/ioutil"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)

	fmt.Print(string(data))
}

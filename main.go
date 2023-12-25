package main

import (
	"aoc/dayx/part1"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	f, err := os.Open("./part1/input.txt")
	checkError(err)

	v, err := part1.Run(f)
	checkError(err)

	fmt.Printf("Your value is %d\n", v)
}

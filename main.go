package main

import (
	d1 "aoc/lib/day1"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Usage: `aoc $day $part`")
		return
	}

	day := args[0]
	part := args[1]

	if day == "1" {
		if part == "1" {
			d1.P1()
		}
	}
}

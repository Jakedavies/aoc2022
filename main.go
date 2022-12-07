package main

import (
	d1 "aoc/lib/day1"
	d2 "aoc/lib/day2"
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

	// uhhh... TODO
	if day == "1" {
		if part == "1" {
			d1.P1()
		}
	}
	if day == "2" {
		if part == "1" {
			d2.P1()
		} else {
			d2.P2()
		}
	}
}

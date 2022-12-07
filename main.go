package main

import (
	d1 "aoc/lib/day1"
	d2 "aoc/lib/day2"
	d3 "aoc/lib/day3"
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
	if day == "3" {
		if part == "1" {
			d3.P1()
		} else {
			d3.P2()
		}
	}
}

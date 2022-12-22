package main

import (
	d1 "aoc/lib/day1"
	d10 "aoc/lib/day10"
	d2 "aoc/lib/day2"
	d3 "aoc/lib/day3"
	d4 "aoc/lib/day4"
	d5 "aoc/lib/day5"
	d6 "aoc/lib/day6"
	d7 "aoc/lib/day7"
	d8 "aoc/lib/day8"
	d9 "aoc/lib/day9"
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

	if day == "4" {
		if part == "1" {
			d4.P1()
		} else {
			d4.P2()
		}
	}

	if day == "5" {
		if part == "1" {
			d5.P1()
		} else {
			d5.P2()
		}
	}

	if day == "6" {
		if part == "1" {
			d6.P1()
		} else {
			d6.P2()
		}
	}

	if day == "7" {
		if part == "1" {
			d7.P1()
		} else {
			d7.P2()
		}
	}

	if day == "8" {
		if part == "1" {
			d8.P1()
		} else {
			d8.P2()
		}
	}
	if day == "9" {
		if part == "1" {
			d9.P1()
		} else {
			d9.P2()
		}
	}
	if day == "10" {
		if part == "1" {
			d10.P1()
		} else {
			d10.P2()
		}
	}

}

package day4

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func P1() {
	path, err := filepath.Abs("./inputs/day4/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	// match number range
	r, err := regexp.Compile(`(\d+)-(\d+),(\d+)-(\d+)`)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		result := r.FindStringSubmatch(line)
		a1, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}
		a2, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}
		b1, err := strconv.Atoi(result[3])
		if err != nil {
			panic(err)
		}
		b2, err := strconv.Atoi(result[4])
		if err != nil {
			panic(err)
		}

		min := Min(a1, b1)
		max := Max(a2, b2)

		if min == a1 && max == a2 {
			count += 1
		} else if min == b1 && max == b2 {
			count += 1
		}
	}

	log.Println("The count of redundant elves is", count)
}

func P2() {
	path, err := filepath.Abs("./inputs/day4/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	// match number range
	r, err := regexp.Compile(`(\d+)-(\d+),(\d+)-(\d+)`)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		result := r.FindStringSubmatch(line)
		a1, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}
		a2, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}
		b1, err := strconv.Atoi(result[3])
		if err != nil {
			panic(err)
		}
		b2, err := strconv.Atoi(result[4])
		if err != nil {
			panic(err)
		}

		min := Min(a1, b1)

		if a1 == min {
			if b1 <= a2 {
				count += 1
			}
		} else if b1 == min {
			if a1 <= b2 {
				count += 1
			}
		}
	}

	log.Println("The count of redundant elves is", count)
}

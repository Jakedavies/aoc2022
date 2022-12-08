package day5

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

func starting_state() [9][]string {
	stacks := [9][]string{
		{"F", "T", "C", "L", "R", "P", "G", "Q"},
		{"N", "Q", "H", "W", "R", "F", "S", "J"},
		{"F", "B", "H", "W", "P", "M", "Q"},
		{"V", "S", "T", "D", "F"},
		{"Q", "L", "D", "w", "V", "F", "Z"},
		{"Z", "C", "L", "S"},
		{"Z", "B", "M", "V", "D", "F"},
		{"T", "J", "B"},
		{"Q", "N", "B", "G", "L", "S", "P", "H"},
	}
	return stacks
}

func P1() {
	path, err := filepath.Abs("./inputs/day5/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	stacks := starting_state()

	// match number range
	r, err := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		if len(line) < 4 || line[:4] != "move" {
			continue
		}
		result := r.FindStringSubmatch(line)

		move_count, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}

		from, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}

		to, err := strconv.Atoi(result[3])
		if err != nil {
			panic(err)
		}

		for i := 0; i < move_count; i++ {
			cur := stacks[from-1]
			if len(cur) > 0 {
				pop := cur[len(stacks[from-1])-1]
				stacks[from-1] = cur[:len(cur)-1]
				stacks[to-1] = append(stacks[to-1], pop)
			}
		}

	}
	log.Println(stacks)
}

func P2() {
	path, err := filepath.Abs("./inputs/day5/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	stacks := starting_state()

	// match number range
	r, err := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		log.Println(stacks)
		line := scanner.Text()
		if len(line) < 4 || line[:4] != "move" {
			continue
		}
		log.Println(line)
		result := r.FindStringSubmatch(line)

		move_count, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}

		from, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}

		to, err := strconv.Atoi(result[3])
		if err != nil {
			panic(err)
		}

		cur := stacks[from-1]
		stack_size := len(cur)
		bottom_of_pickup := stack_size - move_count
		if bottom_of_pickup < 0 {
			bottom_of_pickup = 0
		}

		pop := cur[bottom_of_pickup:]
		stacks[from-1] = cur[:bottom_of_pickup]
		for i := 0; i < len(pop); i++ {
			stacks[to-1] = append(stacks[to-1], pop[i])
		}
	}
	log.Println(stacks)
}

package day2

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func getRoundScore(a string, b string) int {
	switch b {
	case "X":
		mypoints := 1
		if a == "A" {
			// tie
			mypoints += 3
		} else if a == "C" {
			mypoints += 6
		}
		return mypoints
	case "Y":
		mypoints := 2
		if a == "B" {
			// tie
			mypoints += 3
		} else if a == "A" {
			mypoints += 6
		}
		return mypoints
	case "Z":
		mypoints := 3
		if a == "C" {
			// tie
			mypoints += 3
		} else if a == "B" {
			mypoints += 6
		}
		return mypoints
	}
	log.Fatal("invalid input")
	return 0
}

type outcome struct {
	X int
	Y int
	Z int
}

type decision_tree struct {
	A outcome
	B outcome
	C outcome
}

func getScoreForRoundV2(opponent_choice string, expected_outcome string) int {
	d := decision_tree{
		A: outcome{X: 3 + 0, Y: 1 + 3, Z: 2 + 6},
		B: outcome{X: 1 + 0, Y: 2 + 3, Z: 3 + 6},
		C: outcome{X: 2 + 0, Y: 3 + 3, Z: 1 + 6},
	}
	var o outcome
	switch opponent_choice {
	case "A":
		o = d.A
	case "B":
		o = d.B
	case "C":
		o = d.C
	}

	switch expected_outcome {
	case "X":
		return o.X
	case "Y":
		return o.Y
	case "Z":
		return o.Z
	}
	panic("invalid input")
}

func P1() {
	path, err := filepath.Abs("./inputs/day2/2.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		total += getRoundScore(string(line[0]), string(line[2]))
	}

	log.Println("The total score for part 1 is", total)
}

func P2() {
	path, err := filepath.Abs("./inputs/day2/2.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		total += getScoreForRoundV2(string(line[0]), string(line[2]))
	}

	log.Println("The total score for part 2 is", total)
}

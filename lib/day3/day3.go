package day2

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func letter_to_priority(a byte) int {
	ascii := int(a)
	if ascii >= 65 && ascii <= 90 {
		return ascii - 64 + 26
	}

	if ascii >= 97 && ascii <= 122 {
		return ascii - 96
	}

	panic("invalid input")
}

func P1() {
	path, err := filepath.Abs("./inputs/day3/1.txt")
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
		pivot := len(line) / 2

		left := [52]bool{}
		right := [52]bool{}

		for i := 0; i < pivot; i++ {
			v := letter_to_priority(line[i]) - 1
			left[v] = true
		}
		for i := pivot; i < len(line); i++ {
			v := letter_to_priority(line[i]) - 1
			log.Println(v)
			right[v] = true
		}

		for i := 0; i < 52; i++ {
			if left[i] == true && right[i] == true {
				total += i + 1
			}
		}
	}

	log.Println("The total score for part 1 is", total)
}

func P2() {
	// TODO
	path, err := filepath.Abs("./inputs/day3/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		a := [52]bool{}
		b := [52]bool{}
		c := [52]bool{}

		for i := 0; i < len(line1); i++ {
			v := letter_to_priority(line1[i]) - 1
			a[v] = true
		}

		for i := 0; i < len(line2); i++ {
			v := letter_to_priority(line2[i]) - 1
			b[v] = true
		}

		for i := 0; i < len(line3); i++ {
			v := letter_to_priority(line3[i]) - 1
			c[v] = true
		}

		for i := 0; i < 52; i++ {
			if a[i] == true && b[i] == true && c[i] == true {
				total += i + 1
			}
		}
	}

	log.Println("The total score for part 2 is", total)
}

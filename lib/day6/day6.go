package day6

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func P1() {
	path, err := filepath.Abs("./inputs/day6/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	// should only be 1 line?
	line := scanner.Text()
	offset := 4
	count := 0

	keymap := make(map[string]int)
	for i := 0; i < len(line); i++ {
		char := string(line[i])
		// add 1
		if val, ok := keymap[char]; ok {
			keymap[char] = val + 1
		} else {
			keymap[char] = 1
		}
		if keymap[char] == 1 {
			count += 1
		}

		trailing := i - offset
		if trailing >= 0 {
			trailing_char := string(line[trailing])
			keymap[trailing_char] = keymap[trailing_char] - 1
			// remove trailing character from counts
			if keymap[trailing_char] == 0 {
				count -= 1
			}
		}
		if count == offset {
			log.Println("Starting package ends at", i+1)
			break
		}
	}
}

func P2() {
	path, err := filepath.Abs("./inputs/day6/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	// should only be 1 line?
	line := scanner.Text()
	offset := 14
	count := 0

	keymap := make(map[string]int)
	for i := 0; i < len(line); i++ {
		char := string(line[i])
		// add 1
		if val, ok := keymap[char]; ok {
			keymap[char] = val + 1
		} else {
			keymap[char] = 1
		}
		if keymap[char] == 1 {
			count += 1
		}

		trailing := i - offset
		if trailing >= 0 {
			trailing_char := string(line[trailing])
			keymap[trailing_char] = keymap[trailing_char] - 1
			// remove trailing character from counts
			if keymap[trailing_char] == 0 {
				count -= 1
			}
		}
		if count == offset {
			log.Println("Starting package ends at", i+1)
			break
		}
	}
}

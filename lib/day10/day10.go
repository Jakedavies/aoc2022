package day10

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type pair struct {
	x int
	y int
}

func P1() {
	path, err := filepath.Abs("./inputs/day10/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r, err := regexp.Compile(`(\w+) ?(:?-?\d+)?`)

	i := 0
	x := 1
	instruction_counter := 0
	instruction := "noop"
	strength := 0
	registers := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for true {
		if instruction_counter == 0 {
			// execute the pending op
			switch instruction {
			case "addx":
				x += registers[0]
			}
			if !scanner.Scan() {
				break
			}
			line := scanner.Text()
			result := r.FindStringSubmatch(line)

			// queue actions
			switch result[1] {
			case "addx":
				v, err := strconv.Atoi(result[2])
				if err != nil {
					log.Fatal(err)
				}
				instruction = "addx"
				registers = []int{v}
				instruction_counter = 2
			case "noop":
				instruction = "noop"
				registers = nil
				instruction_counter = 1
			}

		}
		instruction_counter -= 1
		i += 1
		log.Println(i, x, x*i)
		if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
			log.Println("strength", x*i)
			strength += x * i
		}
	}
	log.Println("Strength:", strength)

}

func print_screen(width int, height int, screen map[pair]int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if _, ok := screen[pair{x, y}]; ok {
				print("x")
			} else {
				print(" ")
			}
		}
		println()
	}

}

func P2() {
	path, err := filepath.Abs("./inputs/day10/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r, err := regexp.Compile(`(\w+) ?(:?-?\d+)?`)

	i := 0
	width := 40
	height := 6
	sprite_position := 1
	instruction_counter := 0
	instruction := "noop"
	registers := make([]int, 0)
	scanner := bufio.NewScanner(file)
	screen := make(map[pair]int)
	for true {
		if instruction_counter == 0 {
			// execute the pending op
			switch instruction {
			case "addx":
				sprite_position += registers[0]
			}
			if !scanner.Scan() {
				break
			}
			line := scanner.Text()
			result := r.FindStringSubmatch(line)

			// queue actions
			switch result[1] {
			case "addx":
				v, err := strconv.Atoi(result[2])
				if err != nil {
					log.Fatal(err)
				}
				instruction = "addx"
				registers = []int{v}
				instruction_counter = 2
			case "noop":
				instruction = "noop"
				registers = nil
				instruction_counter = 1
			}

		}
		position_in_row := i % width
		row := i / width
		if sprite_position-1 == position_in_row || sprite_position == position_in_row || sprite_position+1 == position_in_row {
			// draw a pixel
			screen[pair{position_in_row, row}] = 1
		}
		// move forward
		instruction_counter -= 1
		i += 1
	}
	print_screen(width, height, screen)
}

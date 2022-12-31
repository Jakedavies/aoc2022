package day12

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

type pair struct {
	x int
	y int
}

type heightMap struct {
	start     pair
	end       pair
	positions map[pair]int
}

func P1() {
	path, err := filepath.Abs("./inputs/day12/2.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	heightMap := heightMap{
		start:     pair{0, 0},
		end:       pair{0, 0},
		positions: make(map[pair]int),
	}

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		// store each character into the heightmap
		for i, char := range line {
			// convert lower case letter to integer based on ascii value
			if char == 'S' {
				heightMap.start = pair{i, y}
			} else if char == 'E' {
				heightMap.end = pair{i, y}
			} else {
				height := int(char) - 96
				heightMap.positions[pair{i, y}] = height
			}
		}
		y += 1
	}

}

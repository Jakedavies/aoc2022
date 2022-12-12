package day8

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type pair struct {
	x int
	y int
}

func P1() {
	path, err := filepath.Abs("./inputs/day8/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// init empty 2d array, unknown size
	var m [][]int = make([][]int, 0)
	// match number range
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, make([]int, 0))
		for i := 0; i < len(line); i++ {
			height, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			m[row] = append(m[row], height)
		}
		row++
	}
	log.Println(m)

	// there is a bug here, if a tree is vis from 2 angles its height is overridden
	visible_map := make(map[pair]int)
	// trace from left
	for row := 0; row < len(m); row++ {
		last_height := -1
		cur_inset := 0
		for cur_inset < len(m[row]) {
			if m[row][cur_inset] > last_height {
				last_height = m[row][cur_inset]
				visible_map[pair{row, cur_inset}] = last_height
			}
			cur_inset++
		}
	}
	// trace from right
	for row := 0; row < len(m); row++ {
		last_height := -1
		cur_inset := len(m[row]) - 1
		for cur_inset > 0 {
			if m[row][cur_inset] > last_height {
				last_height = m[row][cur_inset]
				visible_map[pair{row, cur_inset}] = last_height
			}
			cur_inset--
		}
	}
	// from the top
	for column := 0; column < len(m[0]); column++ {
		last_height := -1
		cur_inset := 0
		for cur_inset < len(m) {
			if m[cur_inset][column] > last_height {
				last_height = m[cur_inset][column]
				visible_map[pair{cur_inset, column}] = last_height
			}
			cur_inset++
		}
	}
	// from the bottom
	for column := 0; column < len(m[0]); column++ {
		last_height := -1
		cur_inset := len(m) - 1
		for cur_inset > 0 {
			if m[cur_inset][column] > last_height {
				last_height = m[cur_inset][column]
				visible_map[pair{cur_inset, column}] = last_height
			}
			cur_inset--
		}
	}

	log.Println("Trees visible from edge", len(visible_map))

	/*
		for row := 0; row < len(m); row++ {
			for column := 0; column < len(m[row]); column++ {
				if visible_map[pair{row, column}] == m[row][column] {
					print("0")
				} else {
					print("1")
				}
			}
			println()
		}
	*/
}

func P2() {
  path, err := filepath.Abs("./inputs/day8/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// init empty 2d array, unknown size
	var m [][]int = make([][]int, 0)
	visible_map := make(map[pair]int)
	// match number range
	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, make([]int, 0))
		for i := 0; i < len(line); i++ {
			height, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			m[row] = append(m[row], height)
		}
		row++
	}
	log.Println(m)
  for i := 0; i < len(m); i++ {
    for j := 0; j < len(m[0]); j++ {
      // trace from right
      last_height := -1
      cur_inset := i
      visible_right := 0
      for cur_inset < len(m[row]) {
        if m[row][cur_inset] > last_height {
          last_height = m[row][cur_inset]
          visible_right++
        }
        cur_inset++
      }
      // trace from left
      last_height = -1
      cur_inset = i
      visible_left := 0
      for cur_inset >= 0 {
        if m[row][cur_inset] > last_height {
          last_height = m[row][cur_inset]
          visible_left++
        }
        cur_inset--
      }
      // from the top
      last_height = -1
      cur_inset = i
      visible_bottom := 0
      for cur_inset < len(m) {
        if m[cur_inset][j] > last_height {
          last_height = m[cur_inset][j]
          visible_bottom++
        }
        cur_inset++
      }

      last_height = -1
      cur_inset = i
      visible_top := 0
      for cur_inset >= 0 {
        if m[cur_inset][j] > last_height {
          last_height = m[cur_inset][j]
          visible_top++
        }
        cur_inset--
      }
      visible_map[pair{i, j}] = visible_bottom * visible_top * visible_right * visible_left
    }
  }
}

package day8

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/exp/maps"
)

type pair struct {
	x int
	y int
}

func trace(m [][]int, origin pair, step pair, max int) map[pair]int {
	cur := origin
	x_max := len(m[0])
	y_max := len(m)
	x_min := 0
	y_min := 0
	cur_max_tree_height := -1
	visible := make(map[pair]int)

	for cur.x < x_max && cur.x >= x_min && cur.y < y_max && cur.y >= y_min {
		current_tree_height := m[cur.y][cur.x]
		if current_tree_height > cur_max_tree_height && current_tree_height <= max {
			visible[cur] = current_tree_height
			cur_max_tree_height = current_tree_height
		}
		cur.x += step.x
		cur.y += step.y
	}

	return visible
}

func tracev2(m [][]int, origin pair, step pair, max int) map[pair]int {
	cur := origin
	x_max := len(m[0])
	y_max := len(m)
	x_min := 0
	y_min := 0
	cur_max_tree_height := -1
	visible := make(map[pair]int)

	for cur.x < x_max && cur.x >= x_min && cur.y < y_max && cur.y >= y_min && cur_max_tree_height < max {
		current_tree_height := m[cur.y][cur.x]
		visible[cur] = current_tree_height
		if current_tree_height > cur_max_tree_height {
			cur_max_tree_height = current_tree_height
		}
		cur.x += step.x
		cur.y += step.y
	}

	return visible
}

func P1() {
	path, err := filepath.Abs("./inputs/day8/test.txt")
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
		m_visible := trace(m, pair{0, row}, pair{1, 0}, 9)
		maps.Copy(visible_map, m_visible)
	}
	// trace from right
	for row := 0; row < len(m); row++ {
		m_visible := trace(m, pair{len(m[0]) - 1, row}, pair{-1, 0}, 9)
		maps.Copy(visible_map, m_visible)
	}

	// frm the top
	for column := 0; column < len(m[0]); column++ {
		m_visible := trace(m, pair{column, 0}, pair{0, 1}, 9)
		maps.Copy(visible_map, m_visible)
	}

	// from the bottom
	for column := 0; column < len(m[0]); column++ {
		m_visible := trace(m, pair{column, len(m) - 1}, pair{0, -1}, 9)
		maps.Copy(visible_map, m_visible)
	}

	for i := 0; i < len(m[0]); i++ {
		for j := 0; j < len(m); j++ {
			if _, ok := visible_map[pair{j, i}]; ok {
				print("1")
			} else {
				print("0")
			}
		}
		println()
	}
	log.Println("Trees visible from edge", len(visible_map))
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

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			r_visible := tracev2(m, pair{i + 1, j}, pair{1, 0}, m[j][i])
			l_visible := tracev2(m, pair{i - 1, j}, pair{-1, 0}, m[j][i])
			d_visible := tracev2(m, pair{i, j + 1}, pair{0, 1}, m[j][i])
			u_visible := tracev2(m, pair{i, j - 1}, pair{0, -1}, m[j][i])
			/*
				log.Println("r_visible", r_visible)
				log.Println("l_visible", l_visible)
				log.Println("d_visible", d_visible)
				log.Println("u_visible", u_visible)
			*/
			visible_map[pair{i, j}] = len(r_visible) * len(l_visible) * len(d_visible) * len(u_visible)
		}
	}

	max := 0
	position := pair{0, 0}

	/*
		for i := 0; i < len(m[0]); i++ {
			for j := 0; j < len(m); j++ {
				print(visible_map[pair{j, i}])
			}
			println()
		}
	*/
	for k, v := range visible_map {
		if v > max {
			position = k
			max = v
		}
	}
	log.Println("Max visible", max, position)
}

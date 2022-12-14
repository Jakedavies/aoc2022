package day9

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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func P1() {
	path, err := filepath.Abs("./inputs/day9/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	r, err := regexp.Compile(`(R|U|D|L) (\d+)`)
	tail_pos := pair{0, 0}
	head_pos := pair{0, 0}

	tail_map := make(map[pair]int)
	tail_map[pair{0, 0}] = 1

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		result := r.FindStringSubmatch(line)
		move_count, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}
		direction := result[1]

		for i := 0; i < move_count; i++ {
			switch direction {
			case "R":
				head_pos.x += 1
			case "L":
				head_pos.x -= 1
			case "D":
				head_pos.y += 1
			case "U":
				head_pos.y -= 1
			}
			delta_x := head_pos.x - tail_pos.x
			delta_y := head_pos.y - tail_pos.y

			abs_x := Abs(delta_x)
			abs_y := Abs(delta_y)

			// if we have more than a 1 position delta
			if abs_x > 1 || abs_y > 1 {
				// clamp the positions
				if delta_x > 1 {
					delta_x = 1
				}
				if delta_x < -1 {
					delta_x = -1
				}
				if delta_y > 1 {
					delta_y = 1
				}
				if delta_y < -1 {
					delta_y = -1
				}

				tail_pos.x += delta_x
				tail_pos.y += delta_y
				tail_map[tail_pos] = 1
				log.Println(head_pos, tail_pos)
			}
		}
	}

	log.Println("tail visited:", len(tail_map))
}

func P2() {
	path, err := filepath.Abs("./inputs/day9/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	r, err := regexp.Compile(`(R|U|D|L) (\d+)`)
	positions := [10]pair{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	tail_map := make(map[pair]int)
	tail_map[pair{0, 0}] = 1

	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		line := scanner.Text()
		result := r.FindStringSubmatch(line)
		move_count, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}
		direction := result[1]

		for i := 0; i < move_count; i++ {
			switch direction {
			case "R":
				positions[0].x += 1
			case "L":
				positions[0].x -= 1
			case "D":
				positions[0].y += 1
			case "U":
				positions[0].y -= 1
			}
			for i := 1; i < 10; i++ {
				leader_pos := positions[i-1]
				delta_y := leader_pos.y - positions[i].y
				delta_x := leader_pos.x - positions[i].x

				abs_x := Abs(delta_x)
				abs_y := Abs(delta_y)

				// if we have more than a 1 position delta
				if abs_x > 1 || abs_y > 1 {
					// clamp the positions
					if delta_x > 1 {
						delta_x = 1
					}
					if delta_x < -1 {
						delta_x = -1
					}
					if delta_y > 1 {
						delta_y = 1
					}
					if delta_y < -1 {
						delta_y = -1
					}

					positions[i].x += delta_x
					positions[i].y += delta_y
					if i == 9 {
						tail_map[positions[i]] = 1
					}
				}
			}
			//log.Println("positions", positions)
		}
	}

	log.Println("tail visited:", len(tail_map))
}

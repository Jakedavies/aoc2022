package day11

import (
	"log"
)

type pair struct {
	x int
	y int
}

type monkey struct {
	items     []int
	new_items []int
	operation func(int) int
	test      func(int) int
	touches   int
}

func P1() {
	monkey_map := make(map[int]*monkey)
	monkey_map[0] = &monkey{[]int{84, 66, 62, 69, 88, 91, 91}, []int{}, func(x int) int { return x * 11 }, func(x int) int {
		if x%2 == 0 {
			return 4
		} else {
			return 7
		}
	}, 0}
	monkey_map[1] = &monkey{[]int{98, 50, 76, 99}, []int{}, func(x int) int { return x * x }, func(x int) int {
		if x%7 == 0 {
			return 3
		} else {
			return 6
		}
	}, 0}
	monkey_map[2] = &monkey{[]int{72, 56, 94}, []int{}, func(x int) int { return x + 1 }, func(x int) int {
		if x%13 == 0 {
			return 4
		} else {
			return 0
		}
	}, 0}
	monkey_map[3] = &monkey{[]int{55, 88, 90, 77, 60, 67}, []int{}, func(x int) int { return x + 2 }, func(x int) int {
		if x%3 == 0 {
			return 6
		} else {
			return 5
		}
	}, 0}
	monkey_map[4] = &monkey{[]int{69, 72, 63, 60, 72, 52, 63, 78}, []int{}, func(x int) int { return x * 13 }, func(x int) int {
		if x%19 == 0 {
			return 1
		} else {
			return 7
		}
	}, 0}
	monkey_map[5] = &monkey{[]int{89, 73}, []int{}, func(x int) int { return x + 5 }, func(x int) int {
		if x%17 == 0 {
			return 2
		} else {
			return 0
		}
	}, 0}
	monkey_map[6] = &monkey{[]int{78, 68, 98, 88, 66}, []int{}, func(x int) int { return x + 6 }, func(x int) int {
		if x%11 == 0 {
			return 2
		} else {
			return 5
		}
	}, 0}
	monkey_map[7] = &monkey{[]int{70}, []int{}, func(x int) int { return x + 7 }, func(x int) int {
		if x%5 == 0 {
			return 1
		} else {
			return 3
		}
	}, 0}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkey_map {
			for j := 0; j < len(monkey.items); j++ {
				monkey.touches += 1
				item := monkey.items[j]
				n := monkey.operation(item)
				new_owner := monkey.test(n)
				if v, ok := monkey_map[new_owner]; ok {
					v.new_items = append(v.new_items, n)
					monkey_map[new_owner] = v
				}
			}
			monkey.items = []int{}
		}

		for _, monkey := range monkey_map {
			monkey.items = monkey.new_items
			monkey.new_items = []int{}
		}
	}
	for i, monkey := range monkey_map {
		log.Printf("Monkey %d touched %d items", i, monkey.touches)
	}
}

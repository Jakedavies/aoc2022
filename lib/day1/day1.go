package day1

import (
	"bufio"
	"container/heap"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Elf struct {
	sum  int
	food []int
}

func P1() {
	path, err := filepath.Abs("./inputs/day1/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var elfs []Elf
	elf_index := 0

	h := &IntHeap{0, 0, 0}
	heap.Init(h)
	// read the file line by line
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		if len(elfs) <= elf_index {
			elfs = append(elfs, Elf{0, []int{}})
		}
		// if line is empty, increment the elf index it to a new elf index, else add it to the sum of the elfs food
		if scanner.Text() == "" {
			elf_index += 1
		} else {
			i, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			// add an item to the food array
			elfs[elf_index].food = append(elfs[elf_index].food, i)
			elfs[elf_index].sum += i

			if (*h)[0] < elfs[elf_index].sum {
				heap.Pop(h)
				heap.Push(h, elfs[elf_index].sum)
			}
		}
	}

	log.Println("The elves with the most food are", h)
	n := 0
	for _, a := range *h {
		n += a
	}
	log.Println("They have a total of ", n)
}

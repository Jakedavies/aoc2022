package day7

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type fs_node interface {
	get_size() int
	get_name() string
}

type directory_node struct {
	name     string
	children map[string]fs_node
	parent   *directory_node
	size     int
}

type file_node struct {
	name string
	size int
}

func (node *directory_node) get_size() int {
	return node.size
}

func (node *directory_node) get_name() string {
	return node.name
}

func (node *file_node) get_size() int {
	return node.size
}

func (node *file_node) get_name() string {
	return node.name
}

func (node *directory_node) add_child(child fs_node) {
	_, exists := node.children[child.get_name()]
	if !exists {
		node.children[child.get_name()] = child
	}
	node.size += child.get_size()
	cur := node
	// bubble up the size
	for cur := cur.parent; cur != nil; cur = cur.parent {
		cur.size += child.get_size()
	}
}

func new_dir(name string, parent *directory_node) *directory_node {
	return &directory_node{name: name, children: make(map[string]fs_node), parent: parent, size: 0}
}

// n * n op... memoizing would make this better
func crawl_tree(cur *directory_node, dirs *[]*directory_node, filter_fn func(int) bool) {
	size := cur.get_size()
	if filter_fn(size) {
		*dirs = append(*dirs, cur)
	}
	for _, child := range cur.children {
		c, ok := child.(*directory_node)
		if ok {
			// craw subdirs
			crawl_tree(c, dirs, filter_fn)
		}
	}
}

func P1() {
	path, err := filepath.Abs("./inputs/day7/1.txt")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// match number range
	var root = new_dir("/", nil)
	var cur *directory_node = root

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	output_mode := false
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if tokens[0] == "$" {
			output_mode = false
			// what command is it
			switch tokens[1] {
			case "cd":
				{
					if tokens[2] == ".." {
						cur = cur.parent
					} else if tokens[2] == "/" {
						cur = root
					} else {
						new_dir_name := tokens[2]
						c, ok := cur.children[new_dir_name].(*directory_node)
						if !ok {
							panic("not a directory")
						}
						cur = c
					}
				}
			case "ls":
				{
					output_mode = true
				}
			}
		} else if output_mode {
			// this must be the output of an ls?
			name := tokens[1]
			if tokens[0] == "dir" {
				cur.add_child(new_dir(name, cur))
			} else {
				size, err := strconv.Atoi(tokens[0])
				if err != nil {
					panic(err)
				}
				cur.add_child(&file_node{name, size})
			}
			// after we ls, check the dir size and add it to the list if its bigly
		} else {
			panic("No command but not in output mode")
		}
	}

	var dirs []*directory_node
	small_filter := func(size int) bool {
		return size <= 100000
	}
	sum := 0
	crawl_tree(root, &dirs, small_filter)
	for _, dir := range dirs {
		sum += dir.get_size()
	}
	log.Println("sum of small dirs is", sum)

	// how do make math.maxint var
	current_min_dir := 1000000000000

	used_space := root.get_size()
	available_space := 70000000 - used_space
	required_space := 30000000 - available_space

	min_size_filter := func(size int) bool {
		return size >= required_space
	}

	dirs = []*directory_node{}
	crawl_tree(root, &dirs, min_size_filter)
	for _, dir := range dirs {
		if dir.get_size() < current_min_dir {
			current_min_dir = dir.get_size()
		}
	}
	log.Println("smallest dir over threshold is", current_min_dir)
}

func P2() {

}

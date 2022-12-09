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
}

type file_node struct {
	name string
	size int
}

func (node *directory_node) get_size() int {
	sum := 0
	for _, child := range node.children {
		sum += child.get_size()
	}
	return sum
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
}

func new_dir(name string, parent *directory_node) *directory_node {
	return &directory_node{name: name, children: make(map[string]fs_node), parent: parent}
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
		log.Println(line)
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
						log.Println("chaning dir to", tokens[2])
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

	log.Println("total fs size is ", root.get_size())
}

func P2() {

}

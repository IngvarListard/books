package main

import "fmt"

const SIZE = 15


type Node struct {
    Value int
	Next *Node
}

type HashTable struct {
    Table map[int]*Node
	Size int
}

func hashFunc(i, size int) int {
    return (i % size)
}

func insert(hash *HashTable, value int) int {
    index := hashFunc(value, hash.Size)
	element := Node{Value: value, Next: hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable) {
    for k := range hash.Table {
		t := hash.Table[k]
		for t != nil {
			fmt.Printf("%d -> ", t.Value)
			t = t.Next
		}
		fmt.Println()
	}
}

func lookup(hash *HashTable, i int) bool {
    index := hashFunc(value, hash.Size)
	if hash.Table[index] != nil {
		t := hash.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{Table: table, Size: SIZE}

	fmt.Println("NUMBER OF SPACES", hash.Size)

	for i := 0; i < 120; i++ {
		insert(hash, i)
	}

	traverse(hash)
}

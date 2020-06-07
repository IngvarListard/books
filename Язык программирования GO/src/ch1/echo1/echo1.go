package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	sep := " "
	for _, a := range os.Args[1:] {
		s += sep + a
	}

	fmt.Println(s)
}

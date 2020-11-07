package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Command line arguments is not provided")
	}

	var max, min float64
	for i := 1; i < len(os.Args); i++ {
		r, _ := strconv.ParseFloat(os.Args[i], 64)

		if r > max {
			max = r
		}

		if r < min {
			min = r
		}
	}

	fmt.Println("MAX:", max)
	fmt.Println("MIN:", min)
}

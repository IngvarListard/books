package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
			os.Exit(1)
		}
	}
	printStats(mem)


	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}
	printStats(mem)
}

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc", mem.Alloc)
	fmt.Println("mem.TotalAlloc", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc", mem.HeapAlloc)
	fmt.Println("mem.NumGC", mem.NumGC)
	fmt.Println("-----")
}

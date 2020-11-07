package main

import (
	"unsafe"
	"fmt"
)

func main() {
	array := [...]int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Print(*pointer, " ")
	memoryAdress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])


	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAdress))
		fmt.Print(*pointer, " ")
		memoryAdress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}

	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAdress))
	fmt.Print("One more:", *pointer, " ")
	memoryAdress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println()
}

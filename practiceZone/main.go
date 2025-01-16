package main

import (
	"fmt"
	"unsafe"
)

func test(v []float32) {
	fmt.Println(unsafe.Pointer(&v[0]))
}

func main() {
	var value float32 = 123
	t := []*float32{&value}
	for x := range t {
		fmt.Println(&x)
		fmt.Println(x)
	}
}

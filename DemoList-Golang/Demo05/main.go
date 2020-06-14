package main

import "fmt"

func zeroValue(value int) {
	value = 0
}

func zeroPtr(intPointer *int) {
	*intPointer = 0
}

func main() {
	i := 1
	fmt.Println("zeroValue: ", i)

	zeroValue(i)
	fmt.Println("zeroValue: ", i)

	zeroPtr(&i)
	fmt.Println("zeroPtr: ", i)

	// 指针也是可以被打印的。
	fmt.Println("pointer:", &i)
}
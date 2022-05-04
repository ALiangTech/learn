package main

import "fmt"

// 指针 允许在程序中通过 引用传递 来传递值和数据结构

func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}
func main() {
	i := 1
	fmt.Println("initial", i)

	zeroval(i)
	fmt.Println("zeroval", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("pointer", &i)
}

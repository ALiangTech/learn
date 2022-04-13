package main

import "fmt"

// 为结构体类型定义方法

type rect struct {
	width, height int
}

// 给rect 定义一个area方法

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 20, height: 20}
	fmt.Println("rect", r.area())

	fmt.Println("perim", r.perim())

	rp := &r

	fmt.Println("area-rp", rp.area())
	fmt.Println("perim-rp", rp.perim())
}

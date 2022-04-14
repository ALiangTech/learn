package main

import "fmt"

type base struct {
	num  int
	name string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num= %v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{num: 1, name: "hanjun"},
		str:  "some name",
	}
	fmt.Printf("coo={num:%v, str:%v, name:%v}", co.num, co.str, co.name)
	fmt.Println("also num", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer", d.describe())
}

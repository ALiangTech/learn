package main

import "fmt"

// 结构体 带类型的字段的集合
// 类似于js中对象的方式 哈哈
// 关键字 type  struct
// 结构体是可变(mutable)的。
type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 18
	return &p
}

func main() {
	fmt.Println(Person{"Bob", 29})

	fmt.Println(Person{"hanuun", 18})

	fmt.Println(Person{age: 23, name: "hh"})

	s := newPerson("hanjun")
	fmt.Println(s.name)
	sp := s
	sp.age = 51
	fmt.Println(sp.age)
}

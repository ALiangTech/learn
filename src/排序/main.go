package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "B", "d"}
	sort.Strings(strs)

	fmt.Println(strs)

	ints := []int{3, 2, 1}
	sort.Ints(ints)

	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println(s)

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

// 自定义排序

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

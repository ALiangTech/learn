package main

import "fmt"

// 字符串是个byte类型的切片
// 字符 rune 是一个unicode的整数
// https://go.dev/blog/strings
// []byte
func main() {
	const s = "สวัสดี"
	fmt.Println("Len", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x \n", s[i]) // Unicode 的所有字节的十六进制值
	}
}

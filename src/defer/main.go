package main

import (
	"fmt"
	"os"
)

// 用于确保程序在执行完成后，会调用某个函数 一半用于执行清理工作
// 类似于js 中的微任务 宏任务
func main() {
	f := createFile("/root/learn/src/defer/defer.js")
	writeFile(f)
	defer closeFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "const age=12 ")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

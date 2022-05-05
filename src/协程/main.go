package main

// 协程 可以理解成轻量化的进程 但是
// 类似于js 中的异步操作
import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go f("test")

	time.Sleep(time.Second)
	f("test2")
}

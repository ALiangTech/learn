package main

import (
	"fmt"
	"time"
)

// 解决多个协程之间的通信问题
// 是一个管理多个协程的管道
// 默认channel 的接收和发送 是阻塞的
// 密集计算 将结果发送给另外的一个协程
func test() {
	fmt.Println("first")
}
func main() {
	message := make(chan string)

	go func() {
		message <- "ping" // 发送
	}()

	msg := <-message // 接收
	test()
	fmt.Println(msg)
	fmt.Printf("通道缓冲\n")
	cacheChan()
	fmt.Printf("通道同步\n")
	done := make(chan bool, 1)
	go worker(done)
	<-done
	fmt.Print("complete")
	fmt.Printf("\n通道方向\n")
	testpp()
	fmt.Printf("\n通道选择器\n")
	chanSelect()
	fmt.Printf("\n非阻塞通道操作\n")
	testNoStop()
	fmt.Printf("\n通道的关闭\n")
	closeChan()
	fmt.Printf("\n通道的遍历\n")
	rangechan()
}

// 通道缓冲
// 默认有发送 必须 有接收
// 如果有缓冲 可以先发送
// 超过缓冲值 会发生死锁
func cacheChan() {
	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"
	// message <- "thrid" // deadlock 死锁

	fmt.Println(<-message)
	fmt.Println(<-message)
}

// 通道同步
// 使用通道能传递信息 来同步协程之间的执行状态
func worker(done chan bool) {
	fmt.Print("working....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

//通道方向
// 当通道作为参数 可以确定通道是只读 还是只写的
// 提高安全

// 写
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 读 写
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func testpp() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// 通道选择器
// 同时等待多个通道操作
func chanSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}

// 非阻塞通道操作
// 默认通道是阻塞的

func testNoStop() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

// 通道的关闭
// 通道关闭 无法咋向通道操作

func closeChan() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

}

// 通道遍历

func rangechan() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}

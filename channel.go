package main

import (
	"fmt"
	"sync"
)

func write(ch chan int, n int) {
	l := new(sync.Mutex)
	l.Lock()
	ch <- n
	fmt.Printf("write: %d\n", n)
	l.Unlock()
}

func main() {
	// 声明
	var data chan int

	// 初始化
	data = make(chan int) // 数据交换队列
	//data = make(chan int, 3) // 带缓冲区的数据交换队列
	exit := make(chan bool) // 退出通知

	go func() {
		// 从队列迭代接收数据，直到 close
		for d := range data {
			fmt.Printf("read: %d\n", d)
		}
		fmt.Println("recv over.")
		// 发出退出通知
		exit <- true
	}()

	// 发送数据
	data <- 1
	fmt.Printf("write: %d\n", 1)
	data <- 2
	fmt.Printf("write: %d\n", 2)

	// 关闭队列
	close(data)
	fmt.Println("send over.")

	// 接收数据
	<-exit
}

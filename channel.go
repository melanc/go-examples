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
	data := make(chan int) // 数据交换队列
	//data := make(chan int, 3) // 带缓冲区的数据交换队列
	exit := make(chan bool) // 退出通知

	go func() {
		for d := range data { // 从队列迭代接收数据，直到 close 。
			fmt.Printf("read: %d\n", d)
		}
		fmt.Println("recv over.")
		exit <- true // 发出退出通知。
	}()

	write(data, 1) // 发送数据。
	write(data, 2)
	write(data, 3)
	close(data) // 关闭队列。
	fmt.Println("send over.")
	<-exit // 等待退出通知。
}

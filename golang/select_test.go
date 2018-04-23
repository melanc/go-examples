package base

import (
	"fmt"
	"os"
	"testing"
)

func TestSelect(t *testing.T) {
	a, b := make(chan int, 3), make(chan int)
	go func() {
		v, ok, s := 0, false, ""
		for {
			// 随机选择可⽤ channel，接收数据
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}
			if ok {
				fmt.Println(s, v)
			} else {
				os.Exit(0)
			}
		}
	}()
	for i := 0; i < 5; i++ {
		// 随机选择可⽤ channel，发送数据
		select {
		case a <- i:
		case b <- i:
		}
	}
	close(a)

	// 没有可⽤ channel，阻塞 main goroutine
	select {}
}

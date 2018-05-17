package golang

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAddx(t *testing.T) {
	num := new(int32)
	*num = 10
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go Addx(wg, num, 500)
	go Addx(wg, num, 600)
	wg.Wait()
	fmt.Printf("num: %d\n", *num)
}

func TestAtomicStore(t *testing.T) {
	num := new(int32)
	*num = 10
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go AtomicStore(wg, num, 500)
	go AtomicStore(wg, num, 600)
	wg.Wait()
	fmt.Printf("num: %d\n", *num)
}

func TestAtomicAdd(t *testing.T) {
	// atomic
	num := new(int32)
	*num = 10
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go AtomicAdd(wg, num, 500)
	go AtomicAdd(wg, num, 600)
	wg.Wait()
	fmt.Printf("num: %d\n", *num)
}

func TestAtomicSwap(t *testing.T) {
	// atomic
	num := new(int32)
	*num = 10
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go AtomicSwap(wg, num, 500)
	go AtomicSwap(wg, num, 600)
	wg.Wait()
	fmt.Printf("num: %d\n", *num)
}

func TestAtomicCompareAndSwap(t *testing.T) {
	// atomic
	num := new(int32)
	*num = 10
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go AtomicCompareAndSwap(wg, num, 500)
	go AtomicCompareAndSwap(wg, num, 600)
	wg.Wait()
	fmt.Printf("num: %d\n", *num)
}

func Addx(wg *sync.WaitGroup, opt *int32, times int) {
	defer wg.Done()
	var i int
	for i = 0; i < times; i++ {
		*opt = *opt + 1
		//*opt += 1 // 原子操作
	}
}

// 不能保证原子性
func AtomicStore(wg *sync.WaitGroup, opt *int32, times int) {
	defer wg.Done()
	var i int
	for i = 0; i < times; i++ {
		atomic.StoreInt32(opt, atomic.LoadInt32(opt)+1)
	}
}

func AtomicAdd(wg *sync.WaitGroup, opt *int32, times int) {
	defer wg.Done()
	var i int
	for i = 0; i < times; i++ {
		atomic.AddInt32(opt, 1)
	}
}

func AtomicSwap(wg *sync.WaitGroup, opt *int32, times int) {
	defer wg.Done()
	var i int
	for i = 0; i < times; i++ {
		atomic.SwapInt32(opt, *opt+1)
	}
}

func AtomicCompareAndSwap(wg *sync.WaitGroup, opt *int32, times int) {
	defer wg.Done()
	var i int
	for i = 0; i < times; i++ {
		atomic.CompareAndSwapInt32(opt, *opt, *opt+1)
	}
}

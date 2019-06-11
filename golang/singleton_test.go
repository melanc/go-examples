package golang

import "sync"

type singleton struct{}

var ins *singleton
var ins2 *singleton = &singleton{}

var mu sync.Mutex
var once sync.Once

// 懒汉模式，非线程安全
func GetIns1() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}

// 饿汉模式
func GetIns2() *singleton {
	return ins2
}

// 懒汉模式，线程安全
func GetIns3() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}

// 饿汉模式，双重检查
func GetIns4() *singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			ins = &singleton{}
		}
	}
	return ins
}

// sync.Once
func GetIns5() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

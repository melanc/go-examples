package golang

//声明一个常量

const MAX = 4096

//声明一个指定类型的常量

const LIMIT int16 = 1024
const LIMIT2 = int16(1024)

//声明一组常量

const (
	start  = 0x1
	resume = 0x2
	stop   = 0x4
)

//声明一组指定类型的常量
const (
	start2  int8 = 0x1
	resume2 int8 = 0x2
	stop2   int8 = 0x4
)

//用iota简化上面的写法
const (
	start3 int8 = 1 << iota
	resume4
	stop4
)

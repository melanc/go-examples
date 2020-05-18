package other

import (
	"fmt"
	"strings"
	"testing"
)

func TestBitmap(t *testing.T) {
	array := [...]uint64{0, 6, 3, 7, 2, 8, 1, 4}

	var maxNum uint64 = 9
	bm := NewBitmap(maxNum)

	for _, v := range array {
		bm.Set(v)
	}
	bm.Set(5)
	fmt.Println(bm.IsFully())
	fmt.Println(bm.IsEmpty())
	fmt.Println("bitmap 中存在的数字:")
	fmt.Println(bm.GetData())
	fmt.Println("bitmap 中的二进制串")
	fmt.Println(bm.String())
	fmt.Println("bitmap 中的数字个数:", bm.Count())
	fmt.Println("bitmap size:", bm.Size())
	fmt.Println("Contains(0):", bm.Contains(0))
	bm.Reset(5)
	fmt.Println(bm.String())
	fmt.Println("Contains(5):", bm.Contains(5))
	fmt.Println(bm.GetData())
}

const (
	bitSize = 8 // 一个字节8位
)

var bitmask = []byte{1, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7}

// 首字母小写 只能调用 工厂函数 创建
type bitmap struct {
	bits     []byte
	bitCount uint64 // 已填入数字的数量
	capacity uint64 // 容量
}

// 创建工厂函数
func NewBitmap(maxnum uint64) *bitmap {
	return &bitmap{bits: make([]byte, (maxnum+7)/bitSize), bitCount: 0, capacity: maxnum}
}

// 填入数字
func (this *bitmap) Set(num uint64) {
	byteIndex, bitPos := this.offset(num)
	// 1 左移 bitPos 位 进行 按位或 (置为 1)
	this.bits[byteIndex] |= bitmask[bitPos]
	this.bitCount++
}

// 清除填入的数字
func (this *bitmap) Reset(num uint64) {
	byteIndex, bitPos := this.offset(num)
	// 重置为空位 (重置为 0)
	this.bits[byteIndex] &= ^bitmask[bitPos]
	this.bitCount--
}

// 数字是否在位图中
func (this *bitmap) Contains(num uint64) bool {
	byteIndex := num / bitSize
	if byteIndex >= uint64(len(this.bits)) {
		return false
	}
	bitPos := num % bitSize
	// 右移 bitPos 位 和 1 进行 按位与
	return !(this.bits[byteIndex]&bitmask[bitPos] == 0)
}

func (this *bitmap) offset(num uint64) (byteIndex uint64, bitPos byte) {
	byteIndex = num / bitSize // 字节索引
	if byteIndex >= uint64(len(this.bits)) {
		panic(fmt.Sprintf(" runtime error: index value %d out of range", byteIndex))
		return
	}
	bitPos = byte(num % bitSize) // bit位置
	return byteIndex, bitPos
}

// 位图的容量
func (this *bitmap) Size() uint64 {
	return uint64(len(this.bits) * bitSize)
}

// 是否空位图
func (this *bitmap) IsEmpty() bool {
	return this.bitCount == 0
}

// 是否已填满
func (this *bitmap) IsFully() bool {
	return this.bitCount == this.capacity
}

// 已填入的数字个数
func (this *bitmap) Count() uint64 {
	return this.bitCount
}

// 获取填入的数字切片
func (this *bitmap) GetData() []uint64 {
	var data []uint64
	count := this.Size()
	for index := uint64(0); index < count; index++ {
		if this.Contains(index) {
			data = append(data, index)
		}
	}
	return data
}

func (this *bitmap) String() string {
	var sb strings.Builder
	for index := len(this.bits) - 1; index >= 0; index-- {
		sb.WriteString(byteToBinaryString(this.bits[index]))
		sb.WriteString(" ")
	}
	return sb.String()
}

func byteToBinaryString(data byte) string {
	var sb strings.Builder
	for index := 0; index < bitSize; index++ {
		if (bitmask[7-index] & data) == 0 {
			sb.WriteString("0")
		} else {
			sb.WriteString("1")
		}
	}
	return sb.String()
}

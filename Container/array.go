package Container

import (
	"bytes"
	"fmt"
)

type ArrayInterface interface {
	// 添加
	Add(int, interface{}) // 插入元素
	AddLast(interface{})
	AddFirst(interface{})
	// 删除
	Remove(int) interface{}
	RemoveFirst() interface{}
	RemoveLast() interface{}
	// 查找
	Find(interface{}) int      // 查找元素返回第一个索引
	FindAll(interface{}) []int // 查找元素返回所有索引
	Contains(interface{}) bool // 查找是否存在元素
	Get(int) interface{}
	// 修改
	Set(int, interface{})
	// 基本方法
	GetCapacity() int // 获得数组容量
	GetSize() int     // 获得元素个数
	IsEmpty() bool    // 查看数组是否为空
	// 自调用的方法
	resize() // 用于数组的扩、缩容
}

type Array struct {
	data []interface{} // 泛型数组
	size int           // 元素数量
}

// 获得自定义数组，参数为数组的初始长度
func GetArray(capacity int) *Array {
	arr := &Array{}
	arr.data = make([]interface{}, capacity)
	arr.size = 0
	return arr
}

// 获得数组容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 获得数组元素个数
func (a *Array) GetSize() int {
	return a.size
}

// 判断数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// newCapacity 新数组容量
// 逻辑：声明新的数组，将原数组的值 copy 到新数组中
func (a *Array) resize(newCapacity int) {
	newArr := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.data[i]
	}
	a.data = newArr
}

// 获得元素的首个索引，不存在则返回 -1
func (a *Array) Find(element interface{}) int {
	for i := 0; i < a.size; i++ {
		if element == a.data[i] {
			return i
		}
	}
	return -1
}

// 获得元素的所有索引，返回索引组成的切片
func (a *Array) FindAll(element interface{}) (indexes []int) {
	for i := 0; i < a.size; i++ {
		if element == a.data[i] {
			indexes = append(indexes, i)
		}
	}
	return
}

// 查看数组是否存在元素，返回 bool
func (a *Array) Contains(element interface{}) bool {
	if a.Find(element) == -1 {
		return false
	}
	return true
}

// 获得索引对应元素，需要判断索引有效范围
func (a *Array) Get(index int) interface{} {
	if index < 0 || index > a.size-1 {
		panic("Get failed, index is illegal.")
	}
	return a.data[index]
}

// 修改索引对应元素值
func (a *Array) Set(index int, element interface{}) {
	if index < 0 || index > a.size-1 {
		panic("Set failed, index is illegal.")
	}
	a.data[index] = element
}

func (a *Array) Add(index int, element interface{}) {
	if index < 0 || index > a.GetCapacity() {
		panic("Add failed, require index >= 0 and index <= capacity")
	}
	// 数组已满则扩容
	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}
	// 将插入的索引位置之后的元素后移，腾出插入位置
	for i := a.size - 1; i > index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = element
	// 维护数组元素的数量
	a.size++
}

func (a *Array) AddLast(element interface{}) {
	a.Add(a.size, element)
}

func (a *Array) AddFirst(element interface{}) {
	a.Add(0, element)
}

func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Remove failed, index is illegal.")
	}

	removeEle := a.data[index]
	// 从 index 之后的元素，都向前移动一个位置
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	// 清理最后一个元素
	a.data[a.size] = nil

	// 考虑边界情况，不能 resize 为0
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.resize(len(a.data) / 2)
	}
	return removeEle
}

func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

func (a *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, a.GetCapacity()))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i != a.size-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

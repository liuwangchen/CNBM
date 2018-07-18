package main

import (
	"CNBM/Container"
	"fmt"
)

func main() {
	arr := Container.GetArray(10)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)
}

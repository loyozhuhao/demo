package main

import "fmt"

func main() {
	b()

	fmt.Println("a return:", aa()) // 打印结果为 a return: 0
}
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func aa() int {
	var i int

	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()

	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()

	return i
}

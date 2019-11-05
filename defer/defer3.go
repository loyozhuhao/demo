package main

import "fmt"

func main() {
	c()
}

func c() (i int) {
	defer func() {
		i++
		fmt.Println("i:", i)
	}()

	return 1
}

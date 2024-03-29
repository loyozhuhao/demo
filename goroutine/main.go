package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	num := runtime.NumCPU()
	fmt.Println(num)
	runtime.GOMAXPROCS(4)

	go say("world")
	say("hello")

	fmt.Println("NumGoroutine", runtime.NumGoroutine())
}

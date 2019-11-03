package main

import "fmt"

func main() {
	for i:=0;i<10;i++{
		 fmt.Printf("i:",i)
	}

	for true  {
		fmt.Printf("这是无限循环。\n");
	}
}

package main

import (
	"fmt"
	"log"
)

type T2 struct{ n int }

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	ts := [2]T2{}
	for i, t := range &ts {
		switch i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n, " ")
		}
	}
	fmt.Print(ts)
}

package main

import (
	"fmt"
	"log"
)

type T3 struct{ n int }

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	ts := [2]T3{}
	for i := range ts[:] {
		t := &ts[i]
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

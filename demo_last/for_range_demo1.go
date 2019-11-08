package main

import (
	"fmt"
	"log"
)

//http://note.youdao.com/noteshare?id=a8176cae082664c88a514c4cb7e7b8fa

type T struct{ n int }

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	ts := [2]T{}
	for i, t := range ts {
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

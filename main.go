package main

import (
	"fmt"
)

func main() {
	list := newSkipList()
	for i := 1; i < 10000; i++ {
		list.insert(i, "b")
	}

	for i := 5000; i > 0; i-- {
		fmt.Println(list.get(i))
	}

}

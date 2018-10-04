package main

import (
	"fmt"

	"esurelabs.homelinux.com/verrol/playground/gophercon/2018/binary-tree/tree"
)

var (
	keys = []rune{'s', 't', 'a', 'c', 'y'}
)

func main() {
	for _, key := range keys {
		t := tree.New()
		for _, k := range keys {
			t.Insert(string(k))
		}

		fmt.Printf("tree orig\n------\n%v\n", t)
		t.Delete(string(key))
		fmt.Printf("tree\n-----\n%v\n", t)
	}

	t := tree.New()
	for _, k := range keys {
		t.Insert(string(k))
	}

	fmt.Printf("min val:\n%s\n", t.Min())
	fmt.Printf("max val:\n%s\n", t.Max())
}

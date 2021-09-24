package main

import (
	"fmt"
	"testing"
)

type node struct {
	i int
}

func TestMain(m *testing.M) {
	c := make(chan *node)
	n := &node{1}
	go func() {
		c <- n
	}()

	//m := <-c
	//m.i = 2
	fmt.Printf("n=%v,%p   m=%v,%p\n", n, n, m, m)
}

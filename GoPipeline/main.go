package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 10

func main() {
	out := random(N)
	print(out)
}

func random(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			x := rand.New(rand.NewSource(time.Now().UTC().UnixNano())).Int()
			out <- x
		}
		close(out)
	}()
	return out
}

func print(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}
package main

import (
	"fmt"
	"iter"
	"sync"
	"time"
)

func Range(size int) iter.Seq[int] {
	return func(yield func(int) bool) {
		wg := sync.WaitGroup{}
		for value := range size {
			wg.Go(func() {
				if !yield(value) {
					return
				}
			})
		}

		wg.Wait()
	}
}

func main() {
	fmt.Println("Started")
	for value := range Range(10) {
		time.Sleep(time.Second * 2)
		fmt.Println(value)
	}
}

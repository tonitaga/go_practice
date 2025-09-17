package main

import (
	"fmt"
	"iter"
	"slices"
	"sync"
	"time"
)

func ConcurrentRange(sequence iter.Seq[int]) {
	wg := sync.WaitGroup{}

	next, stop := iter.Pull(sequence)

	defer stop()
	for {
		value, ok := next()
		if !ok {
			break
		}

		wg.Go(func() {
			time.Sleep(time.Second * 2)
			fmt.Println(value)
		})

	}

	wg.Wait()
}

func main() {
	ConcurrentRange(slices.Values([]int{1, 2, 3, 4, 5, 6}))
}

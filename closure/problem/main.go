package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Go(func ()  {
			fmt.Println(i)
		})
	}

	wg.Wait()

	// До версии 1.22 выводилось всегда 5 для всех запущенных горутин
	// Решение до 1.22

	fmt.Println("Solve-#1-----------------------")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func (v int)  {
			defer wg.Done()
			fmt.Println(v)
		}(i)
	}

	wg.Wait()

	fmt.Println("Solve-#2-----------------------")

	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}

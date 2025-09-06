package main

import (
	"fmt"
	"time"
)

func StopOrDone[T any, Y any](dataChannel <-chan T, stopToken <-chan Y) <-chan T {
	outputChannel := make(chan T)

	go func() {
		defer close(outputChannel)

		for {
			/*
				Дополнительная приоритезация для того,
				чтобы избежать случайного выбора ветки в момент одновременного получения
				сигнала на стоп и чтения значения
			*/
			select {
			case <-stopToken:
				return
			default:
			}

			select {
			case value, opened := <-dataChannel:
				if !opened {
					fmt.Println("Exiting from data end")
					return
				}

				outputChannel <- value
			case <-stopToken:
				fmt.Println("Exiting from stop token")
				return
			}
		}
	}()

	return outputChannel
}

func main() {
	dataChannel := make(chan string)

	go func() {
		for i := 0; i < 50; i++ {
			dataChannel <- fmt.Sprintf("Data #%d", i)
			time.Sleep(time.Millisecond * 100)
		}

		close(dataChannel)
	}()

	done := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		close(done)
	}()

	for value := range StopOrDone(dataChannel, done) {
		fmt.Println(value)
	}
}

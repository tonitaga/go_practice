package main

import (
	"fmt"
	"sync"
)

func ParseData(dataChan <- chan string) <- chan string {
	parsedChan := make(chan string)

	go func ()  {
		defer close(parsedChan)

		index := 0
		for value := range dataChan {
			parsedChan <- fmt.Sprintf("[parsed][%02d] %s", index, value)
			index++
		}
	}()

	return parsedChan 
}

func SendParsedData(parsedData <- chan string, n int) <- chan string {
	outputChan := make(chan string)

	wg := sync.WaitGroup{}
	for i := range n {
		wg.Go(func ()  {
			for value := range parsedData {
				outputChan <- fmt.Sprintf("[send][%02d] %s", i, value)
			}
		})
	}

	go func ()  {
		wg.Wait()
		close(outputChan)
	}()

	return outputChan
}

func main() {
	mainChan := make(chan string)

	go func ()  {
		defer close(mainChan)

		for i := 0; i != 64; i++ {
			mainChan <- fmt.Sprintf("Data from storage #%02d", i)
		}
	}()

	for value := range SendParsedData(ParseData(mainChan), 4) {
		fmt.Println(value)
	}
}
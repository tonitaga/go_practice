package main

import (
	"fmt"
	"sync"
)

func ParseData(dataChan <- chan string, n int) []<- chan string {
	parsedChans := make([]chan string, 0, n)
	for i := 0; i < n; i++ {
		parsedChans = append(parsedChans, make(chan string))
	}

	go func ()  {
		index := 0
		for value := range dataChan {
			parsedChans[index] <- fmt.Sprintf("[parsed] %s", value)
			index = (index + 1) % n
		}

		for _, channel := range parsedChans {
			close(channel)
		}
	}()

	resultChans := make([]<- chan string, n)
	for i, channel := range parsedChans {
		resultChans[i] = channel
	}

	return resultChans 
}

func SendParsedData(parsedData <- chan string) <- chan string {
	outputChan := make(chan string)

	go func ()  {
		defer close(outputChan)

		for value := range parsedData {
			outputChan <- fmt.Sprintf("[send] %s", value)
		}
	}()

	return outputChan
}

func main() {
	mainChan := make(chan string)

	go func ()  {
		defer close(mainChan)

		for i := 0; i != 128; i++ {
			mainChan <- fmt.Sprintf("Data from storage #%02d", i)
		}
	}()

	parsedChans := ParseData(mainChan, 2)
	
	wg := sync.WaitGroup{}
	for i, channel := range parsedChans {
		wg.Go(func ()  {
			for value := range SendParsedData(channel) {
				fmt.Printf("[chan][%d] %s\n", i, value)
			}
		})
	}

	wg.Wait()
}

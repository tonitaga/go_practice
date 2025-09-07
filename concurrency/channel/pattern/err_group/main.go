package main

/*
Errgroup — пакет в языке программирования Go (Golang),
который предоставляет синхронизацию нескольких горутин и сбор их ошибок.
*/

import (
	"fmt"
	"sync"
)

type DoneToken struct{}

type ErrorGroupTask func() error

type ErrorGroup struct {
	err      error
	doneChan chan DoneToken
	wg       sync.WaitGroup
	callOnce sync.Once
}

func NewErrorGroup() (*ErrorGroup, chan DoneToken) {
	doneChan := make(chan DoneToken)

	errorGroup := &ErrorGroup{
		err:      nil,
		doneChan: doneChan,
		wg:       sync.WaitGroup{},
		callOnce: sync.Once{},
	}

	return errorGroup, doneChan
}

func (errorGroup *ErrorGroup) Go(task ErrorGroupTask) {
	errorGroup.wg.Go(func() {
		select {
		case <-errorGroup.doneChan:
			return
		default:
			if err := task(); err != nil {
				errorGroup.callOnce.Do(func() {
					errorGroup.err = err
					close(errorGroup.doneChan)
				})
			}
		}
	})
}

func (errorGroup *ErrorGroup) Wait() error {
	errorGroup.wg.Wait()
	return errorGroup.err
}

func main() {
	errorGroup, _ := NewErrorGroup()

	for i := 0; i != 5; i++ {
		errorGroup.Go(func() error {
			if i == 0 {
				return fmt.Errorf("error occured #%d", i)
			}

			return nil
		})
	}

	if err := errorGroup.Wait(); err != nil {
		fmt.Println(err)
	}
}

package main

import "fmt"

/*

Итератор в Go 1.23 — это функция, которая последовательно проходит
через элементы последовательности и отправляет их в функцию обратного
вызова, обычно называемую yield.

Функция останавливается, когда достигает конца последовательности или
когда yield сигнализирует о раннем прекращении, возвращая false.

*/

func Range(yield func(int) bool) {
	for value := range 10 {
		if !yield(value) {
			return
		}
	}
}

func main() {
	for value := range Range {
		fmt.Println(value)
	}

	fmt.Println("========")

	for value := range Range {
		fmt.Println(value)

		if value == 4 {
			break
		}
	}
}

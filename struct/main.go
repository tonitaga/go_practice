package main

import "fmt"

type Human struct {
	name string
	surname string
	age int
}

func NewHuman(name string, surname string, age int) *Human {
	return &Human{
		name: name,
		surname: surname,
		age: age,
	}
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) GetSurname() string {
	return h.surname
}

func (h *Human) SetSurname(surname string) {
	h.surname = surname
}

func (h *Human) GetAge() int {
	return h.age
}

func (h *Human) SetAge(age int) {
	h.age = age
}

func main() {
	human := NewHuman("Nurislam", "Gubaydullin", 22)
	fmt.Println(*human)
}
package main

import (
	"fmt"
)

type Human struct {
	firstName  string
	secondName string
	age        int
}

type Action struct {
	Human
}

func (a Action) sleep(time int) {
	fmt.Printf("%s пошел спать %d часов\n", a.firstName, time)
}

func (a Action) standUp() {
	fmt.Printf("%s проснулся и готов к работе!\n", a.firstName)
}

func (h Human) print() {
	fmt.Printf("Имя: %s\n", h.firstName)
	fmt.Printf("Фамилия: %s\n", h.secondName)
	fmt.Printf("Возраст: %d\n", h.age)
}

func main() {
	jack := Action{Human: Human{
		"Jack",
		"Lee",
		33,
	}}
	jack.print()
	jack.sleep(10)
	jack.standUp()
}

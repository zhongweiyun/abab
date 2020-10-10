package main

import "fmt"

type Human struct {
	name, phone string
	age	int
}
type Student struct {
	Human
	school string
}
type Employee struct {
	Human
	company string
}
func main() {
	s1 := Student{Human{"Daniel", "15012345***", 13}, "十一中学"}
	e1 := Employee{Human{"Steven", "17812345***", 35}, "1000phone"}
	s1.SayHi()
	e1.SayHi()
}
func (h *Human) SayHi () {
	fmt.Printf("大家好!我是%s ， 我%d岁， 我的联系方式是: %s\n", h.name, h.age ,h.phone)
}

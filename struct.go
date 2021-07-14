package main

import "fmt"

type Person struct {
	age  int
	name string
}

func (p *Person) Sayhello() {
	fmt.Println("Hello :", p.name, " Age :", p.age)
}

func main() {
	var p = new(Person)
	p.name = "Emmanuella"
	p.age = 12
	p.Sayhello()

}

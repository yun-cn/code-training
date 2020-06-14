package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"Bob", 20})

	fmt.Println(Person{name: "Alice", age: 30})

	fmt.Println(Person{name: "Rakuten"})

	fmt.Println(&Person{name: "Google", age: 1})

	p := Person{name: "sean", age: 30}
	fmt.Println(p.name)

	sp := &p
	fmt.Println((*sp).age)

	sp.age = 51
	fmt.Println(sp.age)
}
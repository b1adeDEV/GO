package main

import "fmt"

type person struct {
	name string
	age  int
	year int
}

func createNewPerson(name string) *person {
	p := person{name: name}
	p.age = 18
	p.year = 2007
	return &p
}
func main() {
	a := person{"Bob", 20, 1984}
	b := &a
	b.name = "aidos"
	fmt.Println(b)
	fmt.Println(person{"Bob", 20, 1984})
	fmt.Println(createNewPerson("Blade"))
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

package structs

import "fmt"

type person struct {
	name string
	age  int
}

func StructBased() {
	p := person{}
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)
}

func StructNew() {
	p := new(person)
	p.name = "Lee"
}

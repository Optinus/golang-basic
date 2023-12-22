package struc

import "fmt"

type Person struct {
	name string
	age  int
}

func Struct() {
	var person Person

	person.name = "wildan firmani quraisi"
	person.age = 18

	fmt.Println(person.name)
	fmt.Println(person.age)
}

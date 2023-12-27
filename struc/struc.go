package struc

import "fmt"

// Deklarasi struktur data (struct) bernama Person
type Person struct {
	name string
	age  int
}

func Struct() {
	// Mendeklarasikan variabel person dengan tipe data Person (struct)
	var person Person

	// Mengisi nilai atribut name dan age dari variabel person
	person.name = "Wildan Firmani Quraisi"
	person.age = 18

	// Menampilkan nilai atribut name dan age dari variabel person
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
}

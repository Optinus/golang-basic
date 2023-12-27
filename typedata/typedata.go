package typedata

import (
	"fmt"
)

func TypeData() {
	// Deklarasi dan inisialisasi variabel dengan berbagai tipe data
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	// Menampilkan nilai variabel dengan tipe datanya ke layar
	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
}

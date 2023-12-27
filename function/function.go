package function

import "fmt"

// Example adalah fungsi yang menerima dua parameter string dan mengembalikan hasil penggabungan keduanya
func Example(a string, b string) string {
	return a + b
}

// Function adalah fungsi utama yang memanggil fungsi Example dan menampilkan hasilnya ke layar
func Function() {
	// Memanggil fungsi Example dengan dua parameter dan menampilkan hasilnya ke layar
	fmt.Println(Example("example function :", "parameter end return value"))
}

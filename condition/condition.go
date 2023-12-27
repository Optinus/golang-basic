package condition

import "fmt"

func Condition() {
	// Mendeklarasikan dua variabel a dan b dengan nilai yang sama
	a := 10
	b := 10

	// Menggunakan struktur kontrol if-else untuk menentukan hubungan antara a dan b
	if a < b {
		fmt.Println("a lebih kecil dari b")
	} else {
		fmt.Println("a lebih besar atau sama dengan b")
	}
}

package switc

import "fmt"

func Switch() {
	// Penggunaan switch statement untuk memeriksa nilai days
	days := 1

	switch days {
	case 1:
		fmt.Println("Sabtu") // Jika days == 1, cetak "Sabtu"
	case 2:
		fmt.Println("Ahad") // Jika days == 2, cetak "Ahad"
	default:
		fmt.Println("Hari lain") // Jika days bukan 1 atau 2, cetak "Hari lain"
	}
}

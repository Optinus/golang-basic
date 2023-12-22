package condition

import "fmt"

func Condition() {
	a := 10
	b := 10

	if a < b {
		fmt.Println("a lebih besar dari b")
	} else {
		fmt.Println("a lebih kecil atau sama dengan b")
	}
}

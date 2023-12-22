package function

import "fmt"

func Example(a string, b string) string {
	return a + b
}

func Function() {
	fmt.Println(Example("example function :", "parameter end return value"))
}

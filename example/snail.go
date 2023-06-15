package example

import "fmt"

func Snail() {
	var A, B, V int
	fmt.Scanln(&A, &B, &V)

	day := 0

	for {
		day++
		V -= A
		if V <= 0 {
			break
		}
		V += B
	}

	fmt.Println(day)
}

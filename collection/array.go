package collection

import "fmt"

func ArrayBased() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[2:5])
	fmt.Println(s[1:])
	fmt.Println(s[:3])
}

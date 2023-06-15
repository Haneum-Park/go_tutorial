package main

import "fmt"

func main() {
	a := 2
	b := &a // ? b는 a의 포인터
	*b = 30
	fmt.Println(a, *b)
}

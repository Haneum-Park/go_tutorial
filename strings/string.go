package strings

import (
	"bytes"
	"fmt"
	"strings"
)

func StringBased() {
	s := "Hello, World!"
	println(s)
}

func StringFmt() {
	s := "Hello, World!"
	fmt.Println(s)

	str := fmt.Sprintf("%s, %s!", "Hello", "World")
	fmt.Println(str)
}

func StringByte() {
	var b bytes.Buffer
	b.WriteString("H")
	b.WriteString("e")
	b.WriteString("l")
	b.WriteString("l")
	b.WriteString("o")
	b.WriteString(",")
	b.WriteString(" ")
	b.WriteString("W")
	b.WriteString("o")
	b.WriteString("r")
	b.WriteString("l")
	b.WriteString("d")
	b.WriteString("!")
	fmt.Println(b.String())
}

func StringSlice() {
	s := []string{"Hello", "World!"}
	str := strings.Join(s, ", ")
	fmt.Println(str)
}

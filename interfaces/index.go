package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

// ? Rect struct
type Rect struct {
	width, height float64
}

// ? Circle struct
type Circle struct {
	radius float64
}

// ? Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64      { return r.width * r.height }
func (r Rect) perimeter() float64 { return 2 * (r.width + r.height) }

// ? Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64      { return math.Pi * c.radius * c.radius }
func (c Circle) perimeter() float64 { return 2 * math.Pi * c.radius }

func InterfaceBased() {
	r := Rect{10., 20.}
	rArea := r.area()
	rPerim := r.perimeter()

	fmt.Println("rArea: ", rArea)
	fmt.Println("rPerim: ", rPerim)

	c := Circle{10}
	cArea := c.area()
	cPerim := c.perimeter()

	fmt.Println("cArea: ", cArea)
	fmt.Println("cPerim: ", cPerim)
}

func InterfaceShowArea() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() // ? 인터페이스 메서드 호출
		println(a)
	}
}

// func Marshal(v interface{}) ([]byte, error)
// func Println(a ...interface{}) (n int, err error)

func printIt(v interface{}) {
	fmt.Println(v)
}

func Main() {
	var x interface{}
	x = 1
	x = "Tom"

	printIt(x)
}

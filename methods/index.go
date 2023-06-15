package methods

// ? Rect struct
type Rect struct {
	width, height int
}

// ? Rect의 area() 메서드
func (r Rect) area() int {
	return r.width * r.height
}

func MethodBased() {
	rect := Rect{10, 20}
	area := rect.area() // ? 메서드 호출
	println(area)
}

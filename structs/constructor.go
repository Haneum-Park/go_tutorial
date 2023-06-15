package structs

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d // ? 포인터 전달
}

func StructConstructor() {
	dic := newDict()
	dic.data[1] = "A"
	println(dic.data[1])
}

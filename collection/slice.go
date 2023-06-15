package collection

import "fmt"

func SliceBased() {
	s := []int{1, 2, 3} // 슬라이스 변수 선언
	s[1] = 10
	fmt.Println(s) // [1 10 3] 출력
}

func SliceMake() {
	s := make([]int, 5, 10)
	sMerge := fmt.Sprintf("길이 : %d, 용량 : %d\n", len(s), cap(s))
	fmt.Println(sMerge)
}

func SliceAppend() {
	// NOTE 슬라이스는 동적으로 크기가 늘어나며 고정되지 않을 수 있다.
	// NOTE 추가 시 용량 내 길이를 초과하면 자동으로 용량이 현재 용량의 2배에 해당하는 새로운 슬라이스를 생성한다.
	// NOTE 용량이 초과되지 않으면 용량 내 길이만 변경하여 기존 슬라이스를 그대로 사용한다.
	sliceA := make([]int, 0, 3)
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}
	fmt.Println(sliceA)
}

func SliceAppendCopy() {
	s := []int{0, 1}
	s = append(s, 2)
	s = append(s, 3, 4)
	fmt.Println(s)

	s2 := []int{5, 6, 7}
	target := make([]int, len(s2), cap(s2)*2)
	copy(target, s2)

	fmt.Println(target)
	fmt.Println(s2)
	fmt.Println(len(target), cap(target))
	fmt.Println(len(s2), cap(s2))
}

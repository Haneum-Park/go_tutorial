package example

import "fmt"

// ? i         1 2 3 4 5 6 7 8 9
// ? cnt       1 3 5 7 9 7 5 3 1
// ? spaceCnt  4 3 2 1 0 1 2 3 4
// ? tcnt      9 9 9 9 9 9 9 9 9
func Stars1() {
	var n int
	fmt.Scanln(&n)
	tcnt := 2*n - 1

	for i := 1; i <= tcnt; i++ {
		var star string
		cnt := 2*(tcnt-i) + 1
		if i <= tcnt/2 {
			cnt = 2*i - 1
		}

		var spaceCnt = (tcnt - cnt) / 2
		var space string
		for j := 0; j < spaceCnt; j++ {
			space += " "
		}
		for j := 0; j < cnt; j++ {
			star += "*"
		}

		line := fmt.Sprintf("%s%s%s", space, star, space)
		fmt.Println(line)
	}
}

package example

import "fmt"

func ConvertBased() {
	var x float32 = 5.7
	var y int = int(x + 0.5)

	fmt.Println(y)
}

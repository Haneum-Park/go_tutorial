package example

// NOTE 문제
// NOTE 준규가 가지고 있는 동전은 총 N종류이고, 각각의 동전을 매우 많이 가지고 있다.
// NOTE 동전을 적절히 사용해서 그 가치의 합을 K로 만들려고 한다.
// NOTE 이때 필요한 동전 개수의 최솟값을 구하는 프로그램을 작성하시오.

// NOTE 입력
// NOTE 첫째 줄에 N과 K가 주어진다. (1 ≤ N ≤ 10, 1 ≤ K ≤ 100,000,000)
// NOTE 둘째 줄부터 N개의 줄에 동전의 가치 Ai가 오름차순으로 주어진다.
// NOTE (1 ≤ Ai ≤ 1,000,000, A1 = 1, i ≥ 2인 경우에 Ai는 Ai-1의 배수)

// NOTE 출력
// NOTE 첫째 줄에 K원을 만드는데 필요한 동전 개수의 최솟값을 출력한다.

import (
	"fmt"
	"sort"
)

func Coin() {
	var N, K int
	fmt.Scanln(&N, &K)

	coins := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&coins[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	cnt := 0
	for _, i := range coins {
		cnt += K / i
		K %= i
		if K == 0 {
			break
		}
	}

	fmt.Println(cnt)
}

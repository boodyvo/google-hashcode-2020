package computing

import (
	"fmt"
	"math/rand"
	"time"
)

func increase(m, n int, types *[]int, used *[]bool, curResult *int) {
	for i := rand.Intn(n); i >= 0; i-- {
		if !(*used)[i] {
			(*used)[i] = true
			*curResult += (*types)[i]

			return
		}
	}
}

func decrease(m, n int, types *[]int, used *[]bool, curResult *int) {
	tmp := rand.Intn(n)
	for i := tmp; i < n; i++ {
		if (*used)[i] {
			(*used)[i] = false
			*curResult -= (*types)[i]

			return
		}
	}

	for i := tmp - 1; i >= 0; i-- {
		if (*used)[i] {
			(*used)[i] = false
			*curResult -= (*types)[i]

			return
		}
	}
}

func Computing(m, n int, types []int, printf func(f string, a ...interface{})) {
	fmt.Println(m, n)
	fmt.Println(types)
	rand.Seed(time.Now().UnixNano())
	best := 0
	usd := make([]bool, n)

LOOP:
	for it := 0; it < 100; it++ {
		used := make([]bool, n)
		curResult := 0
		for count := 0; count < 10000; count++ {
			if curResult == m {
				fmt.Println(curResult)
				fmt.Println(used)

				break LOOP
			}

			if curResult > m {
				decrease(m, n, &types, &used, &curResult)
			} else {
				increase(m, n, &types, &used, &curResult)
			}

			if best < curResult && curResult < m {
				best = curResult
				copy(usd, used)
			}
		}

		if best < curResult && curResult < m {
			best = curResult
			copy(usd, used)
		}
	}

	fmt.Println(best)
	fmt.Println(usd)
	res := make([]int, 0, n)
	checker := 0
	for i := 0; i < n; i++ {
		if usd[i] {
			checker += types[i]
			res = append(res, i)
		}
	}
	fmt.Println(checker)

	ans := len(res)
	printf("%d\n", ans)
	for i := 0; i < ans; i++ {
		if i != ans - 1 {
			printf("%d ", res[i])
		} else {
			printf("%d\n", res[i])
		}
	}

	return
}
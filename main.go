package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	if n == 1 {
		fmt.Println(1)
	} else {
		as := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&as[i])
		}
		sort.Ints(as)
		fmt.Println(bestCount(as, n-1, m))
	}
}

func bestCount(as []int, max, m int) int {
	maxCount := 0
	right := 0

	for left := 0; left < max; left++ {
		for right < max && as[right] < as[left]+m {
			right++
		}

		var count int
		if as[right] < as[left]+m {
			count = right - left + 1
		} else {
			count = right - left
		}

		maxCount = func(count1, count2 int) int {
			if count1 > count2 {
				return count1
			} else {
				return count2
			}
		}(maxCount, count)
	}

	return maxCount
}

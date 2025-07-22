// https://atcoder.jp/contests/adt_medium_20250522_1/tasks/abc356_c
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"math"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, M, K int
	fmt.Fscan(reader, &N, &M, &K)

	rs := make(map[int]bool)
	ds := make(map[int]bool)
	for i := 0; i < M; i++ {
		var c int
		fmt.Fscan(reader, &c)
	
		as := make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Fscan(reader, &as[j])
		}

		var r string
		fmt.Fscan(reader, &r)

		if r == "o" {
			calc(0, K, as, rs, []int{})
		} else {
			calc(0, K, as, ds, []int{})
		}
	}

	r := 0
	for v, _ := range rs {
		if ds[v] {
			continue
		}

		r = r + 1
	}
	fmt.Println(r)
}

func calc(index int, K int, as []int, m map[int]bool, ns []int) {
	if len(ns) == K {
		m[fix(ns)] = true
	} else {
		for i := index; i < len(as); i++ {
			calc(i + 1, K, as, m, append(ns, as[i]))
		}
	}
}

func fix(ns []int) int {
	sort.Ints(ns)

	r := 0
	for i := 0; i < len(ns); i++ {
		r = r + (ns[i] * int(math.Pow(10, float64(i * 3))))
	}

	return r
}

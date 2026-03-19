// https://atcoder.jp/contests/adt_all_20260319_1/tasks/abc358_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)

	ni := make([][]int, M)
	for i := 0; i < N; i++ {
		var s string
		fmt.Fscan(in, &s)

		chars := strings.Split(s, "")
		for j := 0; j < M; j++ {
			if chars[j] == "o" {
				if ni[j] == nil {
					ni[j] = []int{}
				}

				ni[j] = append(ni[j], i)
			}
		}
	}

	sort.Slice(ni, func(i, j int) bool {
		return len(ni[i]) < len(ni[j])
	})

	idx := 0
	ans := 0
	idxs := []int{}

	nb := make([]bool, N)
	for idx < len(ni) {
		if len(ni[idx]) == 1 {
			if !nb[ni[idx][0]] {
				nb[ni[idx][0]] = true
				ans++
			}
		} else {
			exist := false
			for i := 0; i < len(ni[idx]); i++ {
				if nb[ni[idx][i]] {
					exist = true
					break
				}
			}

			if !exist {
				idxs = append(idxs, idx)
			}
		}

		idx++
	}

	fmt.Fprintln(out, all(ni, 0, ans, idxs, nb))
}

func all(ni [][]int, idx, ans int, idxs []int, nb []bool) int {
	if idx == len(idxs) {
		return ans
	}

	result := -1

	for _, n := range ni[idxs[idx]] {
		b := nb[n]

		if !b {
			nb[n] = true
			ans++
		}

		idx = idx + 1
		t := all(ni, idx, ans, idxs, nb)
		if result == -1 || result > t {
			result = t
		}
		idx = idx - 1

		if !b {
			nb[n] = false
			ans--
		}
	}

	return result
}

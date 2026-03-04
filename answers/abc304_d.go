// https://atcoder.jp/contests/adt_hard_20251217_1/tasks/abc304_d

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Key struct {
    x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var W, H, N int
	fmt.Fscan(in, &W, &H, &N)

	pqs := make([][2]int, N)
	for i := 0; i < N; i++ {
		pqs[i] = [2]int{}
		fmt.Fscan(in, &pqs[i][0], &pqs[i][1])
	}

	var A int
	fmt.Fscan(in, &A)
	as := make([]int, A)
	for i := 0; i < A; i++ {
		fmt.Fscan(in, &as[i])
	}

	var B int
	fmt.Fscan(in, &B)
	bs := make([]int, B)
	for i := 0; i < B; i++ {
		fmt.Fscan(in, &bs[i])
	}

  cnt := make(map[Key]int, N)
  for i := 0; i < N; i++ {
    xi := sort.SearchInts(as, pqs[i][0])
    yi := sort.SearchInts(bs, pqs[i][1])
    cnt[Key{xi, yi}]++
  }

  min := -1
  max := 1
  for _, v := range cnt {
    if min == -1 || min > v {
      min = v
    }

    if max < v {
      max = v
    }
  }

  if len(cnt) != (A+1) * (B+1) {
    min = 0
  }

	fmt.Fprintln(out, min, max)
}

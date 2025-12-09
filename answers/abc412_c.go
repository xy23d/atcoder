// https://atcoder.jp/contests/adt_hard_20251125_1/tasks/abc412_c
// 問題文
// 1 から
// N までの番号がついた
// N 個のドミノがあります。ドミノ
// i の大きさは
// S
// i
// ​
//   です。
// いくつかのドミノを左右一列に並べたあとにドミノを倒すことを考えます。ドミノ
// i が右に向けて倒れる時、ドミノ
// i のすぐ右に置かれているドミノの大きさが
// 2S
// i
// ​
//   以下ならばそのドミノも右に向けて倒れます。

// あなたは
// 2 個以上のドミノを選んで左右一列に並べることにしました。ただし、ドミノの並べ方は次の条件を満たす必要があります。

// 一番左のドミノはドミノ
// 1 である。
// 一番右のドミノはドミノ
// N である。
// ドミノ
// 1 のみを右に向けて倒した時に、最終的にドミノ
// N も右に向けて倒れる。
// 条件を満たすドミノの並べ方は存在しますか？また、存在する場合は最小で何個のドミノを並べる必要がありますか？

// T 個のテストケースが与えられるので、それぞれについて問題を解いてください。

// 制約
// 1≤T≤10
// 5

// 2≤N≤2×10
// 5

// 1≤S
// i
// ​
//  ≤10
// 9

// 全てのテストケースに対する
// N の総和は
// 2×10
// 5
//   以下
// 入力される値は全て整数

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)

	for i := 0; i < T; i++ {
		execute(in, out)
	}
}

func execute(in *bufio.Reader, out *bufio.Writer) {
	var N int
	fmt.Fscan(in, &N)

	ss := make([]int64, N)
	toNs := make([][]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &ss[i])

		toNs[i] = []int{}

		for j := 0; j < i; j ++ {
			if ss[j] * 2 >= ss[i] {
				toNs[j] = append(toNs[j], i)
			}

			if ss[i] * 2 >= ss[j] {
				toNs[i] = append(toNs[i], j)
			}
		}
	}

	mins := make([]int, N)
	ns := []int{}
	for _, v := range toNs[0] {
		if v == N - 1 {
			fmt.Fprintln(out, 2)
			return
		}

		mins[v] = 2
		ns = append(ns, v)
	}

	for len(ns) > 0 {
		tns := []int{}
		a := make([]bool, N)

		for _, i := range ns {
			n := mins[i]
			for _, j := range toNs[i] {
				if a[j] {
					continue
				}

				a[j] = true
				if j == N-1 {
					fmt.Fprintln(out, n+1)
					return
				}

				if mins[j] == 0 || mins[j] > n+1 {
					mins[j] = n + 1
					tns = append(tns, j)
				}
			}
		}

		ns = tns
	}

	fmt.Fprintln(out, -1)
}

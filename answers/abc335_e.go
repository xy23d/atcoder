// https://atcoder.jp/contests/adt_hard_20251211_1/tasks/abc335_e

// 問題文
// N 頂点 
// M 辺の連結な無向グラフがあり、 
// i 番目の辺は頂点 
// U 
// i
// ​
//   と頂点 
// V 
// i
// ​
//   を双方向に結びます。
// また、全ての頂点に整数が書いてあり、頂点 
// v には整数 
// A 
// v
// ​
//   が書かれています。

// 頂点 
// 1 から頂点 
// N への単純なパス ( 同じ頂点を複数回通らないパス ) に対して、以下のように得点を定めます。

// パス上の頂点に書かれた整数を通った順に並べた数列 を 
// S とする。
// S が広義単調増加になっていない場合、そのパスの得点は 
// 0 である。
// そうでない場合、 
// S に含まれる整数の種類数が得点となる。
// 頂点 
// 1 から頂点 
// N への全ての単純なパスのうち、最も得点が高いものを求めてその得点を出力してください。

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

	var N, M int
	fmt.Fscan(in, &N, &M)

	as := make([]int, N)
	ok := make([][]bool, N)
	a := make([][]bool, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &as[i])
		ok[i] = make([]bool, N)
		a[i] = make([]bool, N)
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if as[i] <= as[j] {
				ok[i][j] = true
			}

			if as[j] <= as[i] {
				ok[j][i] = true
			}
		}
	}

	uv := make([][]int, N)
	for i := 0; i < M; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u = u - 1
		v = v - 1

		if ok[u][v] {
			if uv[u] == nil {
				uv[u] = []int{}
			}

			if !a[u][v] {
				a[u][v] = true
				uv[u] = append(uv[u], v)
			}
		}

		if ok[v][u] {
			if uv[v] == nil {
				uv[v] = []int{}
			}

			if !a[v][u] {
				a[v][u] = true
				uv[v] = append(uv[v], u)
			}
		}
	}

	ans := 0
	a2 := make([]bool, N)
	a2[0] = true
	for _, v := range uv[0] {
		t := execute(v, N, as, uv, a2)
		if ans < t {
			ans = t
		}
	}

	fmt.Fprintln(out, ans)
}

func execute(u, N int, as []int, uv [][]int, a []bool) int {
	a[u] = true

	ans := 0
	for _, v := range uv[u] {
		if a[v] {
			continue
		}

		if v == N - 1 {
			point := getPoint(N, a, as)
			if ans < point {
				ans = point
			}

			continue
		}

		a[v] = true
		t := execute(v, N, as, uv, a)
		a[v] = false

		if ans < t {
			ans = t
		}
	}

	return ans
}

func getPoint(N int, a []bool, as []int) int {
	ans := 0
	ta := make(map[int]bool)
	for i := 0; i < N; i++ {
		if a[i] {
			n := as[i]

			if !ta[n] {
				ta[n] = true
				ans++
			}
		}
	}

	return ans
}

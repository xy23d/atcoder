// https://atcoder.jp/contests/abc399/tasks/abc399_c

package main

import (
	"fmt"
)

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{make([]int, n+1)}
	for i := 0; i <= n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	xRoot := uf.Find(x)
	yRoot := uf.Find(y)

	if xRoot == yRoot {
		return false
	}

	if xRoot < yRoot {
		uf.parent[yRoot] = xRoot
	} else {
		uf.parent[xRoot] = yRoot
	}

	return true
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	uf := NewUnionFind(N)

	result := 0

	for i := 0; i < M; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		if !uf.Union(u, v) {
			result++
		}

		fmt.Println(uf)
	}

	fmt.Println(result)
}

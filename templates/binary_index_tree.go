package main

import (
	"fmt"
)

type BIT struct {
	n int
	bit []int
}

func NewBIT(n int) *BIT {
	return &BIT{
		n:   n,
		bit: make([]int, n+1),
	}
}

func (b *BIT) Add(i int, v int) {
	if i <= 0 || i > b.n {
		return
	}
	for i <= b.n {
		b.bit[i] += v
		i += i & -i
	}
}

func (b *BIT) Sum(i int) int {
	if i <= 0 {
		return 0
	}
	if i > b.n {
		i = b.n
	}
	
	sum := 0
	for i > 0 {
		sum += b.bit[i]
		i -= i & -i
	}
	return sum
}

func (b *BIT) RangeSum(l, r int) int {
	if l > r {
		return 0
	}
	return b.Sum(r) - b.Sum(l-1)
}

// データ量が多い場合には経路圧縮も検討
// 経路圧縮: 全てのデータを格納して、sortして、indexを付与、index:valueで変換
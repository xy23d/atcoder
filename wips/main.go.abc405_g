package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

  mod := 998244353

	var N, Q int
	fmt.Fscan(reader, &N, &Q)

  as := make([]int, N)
  for i := 0; i < N; i++ {
    fmt.Fscan(reader, &as[i])
  }

  for i := 0; i < Q; i++ {
    var L, R, X int
    fmt.Fscan(reader, &L, &R, &X)

    asT := as[L-1:R]

    ns := []int{}
    for j := 0; j < len(asT); j++ {
      if asT[j] >= X {
        continue
      }

      ns = append(ns, asT[j])
    }

    fmt.Println(ns)

    var r int
    if len(ns) == 0 {
      r = 1
    } else {
      r = len(ns) % mod
      for j := 1; j < len(ns); j++ {
        r = (r * (len(ns) - j)) % mod
      }
    }
    fmt.Println(r)
  }
}

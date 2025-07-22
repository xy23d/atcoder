package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscan(reader, &N, &M)

  uv := make(map[int][]int)
  vs := []int{}
	for i := 0; i < M; i++ {
    var u, v int
    fmt.Fscan(reader, &u, &v)

    var l int
    var h int
    if u > v {
      l = v
      h = u
    } else {
      h = v
      l = u
    }

    if uv[l] == nil {
       uv[l] =[]int{}
    }
    uv[l] = append(uv[l], h) 

    vs = append(vs, v)
  }

  for i := 0; i < len(vs) - 1; i++ {
    var l int
    var h int
    if vs[i] > vs[i+1] {
      l = vs[i+1]
      h = vs[i]
    } else {
      h = vs[i+1]
      l = vs[i]
    }

    if l == h {
      continue
    }

    r := false
    for _, v := range uv[l] {
      if v == h {
        r = true
        break
      }
    }

    if !r {
      fmt.Println("No")
      os.Exit(0)
    }
  }

  for i := 0; i < M; i++ {
    for j := i + 2; j < M; j++ {
      for _, v := range uv[i] {
        if v == j {
          fmt.Println("No")
          os.Exit(0)
        }
      }
    }
  }

  fmt.Println("Yes")
  os.Exit(0)
}

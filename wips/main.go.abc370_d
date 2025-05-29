package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var H, W, Q int
	fmt.Fscan(reader, &H, &W, &Q)

	t := make(map[int]map[int]bool)

	for j := 0; j < Q; j++ {
		var r, c int
		fmt.Fscan(reader, &r, &c)

		if t[r] == nil {
			t[r] = make(map[int]bool)
		}
		if !t[r][c] {
			t[r][c] = true
		} else {
			i := 1
			for {
        ps := [][]int{{r+1, c},{r-1, c}, {r, c+1}, {r, c-1}}

        fmt.Println(ps)
				result := []
        count := 0

        for _, vs := range ps {
          tr := vs[0]
          tc := vs[1]

          if tr > 0 && tr <= H && tc > 0 && tc <= W {
            count++

            if t[tr] == nil {
              t[tr] = make(map[int]bool)
            }

            if !t[tr][tc] {
              t[tr][tc] = true
              result = true
            }
          }
        }

        if count == 0 {
          break
        }

        if result {
          break
        }

        i++
			}
		}

		fmt.Println(t)
	}

	count := 0
	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			if !t[i][j] {
				count++
			}
		}
	}

	fmt.Println(count)
}

// https://atcoder.jp/contests/adt_medium_20250410_1/tasks/abc341_c
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var H, W, N int
	fmt.Fscan(reader, &H, &W, &N)

	var T string
	fmt.Fscan(reader, &T)
	ts := strings.Split(T, "")
	mvs := [][]int{}
	mi := 0
	mj := 0
	a := make(map[int]map[int]bool)
	for i := 0; i < len(ts); i++ {
		if ts[i] == "L" {
			mj = mj - 1
		}

		if ts[i] == "R" {
			mj = mj + 1
		}

		if ts[i] == "U" {
			mi = mi - 1
		}

		if ts[i] == "D" {
			mi = mi + 1
		}

		if !a[mi][mj] {
			mvs = append(mvs, []int{mi, mj})
			if a[mi] == nil {
				a[mi] = make(map[int]bool)
			}
			a[mi][mj] = true
		}
	}

	ss := make(map[int]map[int]bool)
	rs := make([]map[int][]int)

	for i := 1; i <= H; i++ {
		var s string
		fmt.Fscan(reader, &s)
		chars := strings.Split(s, "")

		for j := 1; j <= W; j++ {
			if ss[i] == nil {
				ss[i] = make(map[int]bool)
			}

			ss[i][j] = chars[j-1] == "#"
			if ss[i][j] {
 	      trs := make([]map[int][]int)

        for key, map := range rs {
          r := true
          if len(map[i]) > 0 {
            for _, v := range map[i] {
              if v == j {
                r = false
              }
            }
          }

          if r {
            trs = append(trs, map[i])
          }
        }

        rs = trs
      } else {
				t := make(map[int][]int)

				r := true
				for _, mv := range mvs {
					if mv[0] <= 0 && mv[1] <= 0 {
						if ss[i+mv[0]][j+mv[1]] {
							r = false
							break
						}
					} else {
						if t[i+mv[0]] == nil {
							t[i+mv[0]] = []int{}
						}

						t[i+mv] = append(t[i+mv[0]], j+mv[1])
					}
				}

				if r {
					rs = append(rs, t)
				}
			}
		}
	}

	result := 0
	for _, vs := range rs {
		r := true

		if r {
			result++
		}
	}

	fmt.Println(result)
}

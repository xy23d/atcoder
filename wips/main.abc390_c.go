// https://atcoder.jp/contests/adt_medium_20250610_1/tasks/abc390_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(reader, &H, &W)

	min := -1
	max := -1
	kMin := -1
	kMax := -1
	r := false
	skipCount := 0
	for i := 0; i < H; i++ {
		tMin, tMax, tKMin, tKMax := check(reader)

		if tKMin == -1 {
			if !r {
				continue
			}

			skipCount++
			continue
		}
		
		if skipCount > 0 {
			fmt.Println("No")
			os.Exit(0)
		}

		if min == -1 {
			min = tMin
		}

		if max == -1 {
			max = tMax
		}

		if kMin == -1 {
			kMin = tKMin
		}

		if kMax == -1 {
			kMax = tKMax
		}

		if min < tMin {
			if min < tKMin {
				fmt.Println("No")
				os.Exit(0)
			}
		} else {
			if kMin <= tMin {
				min = tMin
			} else {
				fmt.Println("No")
				os.Exit(0)
			}
		}

		if max > tMax {
			if max > tKMax {
				fmt.Println("No")
				os.Exit(0)
			}
		} else {
			if kMax >= tMax {
				max = tMax
			} else {
				fmt.Println("No")
				os.Exit(0)
			}
		}

		if kMin < tKMin {
			kMin = tKMin
		}
		if kMax > tKMax {
			kMax = tKMax
		}

		r = true
	}

	fmt.Println("Yes")
}

// min、max。kMin、kMaxを引数でうけて、幅を更新して返す
// skipだった場合はfalse
func check(reader *bufio.Reader) (int, int, int, int) {
	b := false
	w := false
	min := -1
	max := -1
	kMin := -1
	kMax := -1
	var s string
	fmt.Fscan(reader, &s)
	chars := strings.Split(s, "")
	for j := 0; j < len(chars); j++ {
		if chars[j] == "#" {
			if b && w {
				fmt.Println("No")
				os.Exit(0)
			}

			if min == -1 {
				min = j

				if kMin == -1 {
					kMin = j
				}
			}

			kMax = j
			max = j

			b = true
		}

		if chars[j] == "." {
			if b {
				w = true
			} else {
				kMin = -1
			}
		}

		if chars[j] == "?" {
			if kMin == -1 {
				kMin = j
			}

			if !w {
				kMax = j
			}
		}
	}

	if kMin == -1 {
		kMin = min
	}

	if kMax == -1 {
		kMax = max
	}

	return min, max, kMin, kMax
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(reader, &H, &W)

	Ss := make(map[int]string)
	for i := 0; i < H; i++ {
		var s string
		fmt.Fscan(reader, &s)
		Ss[i] = s
	}

	snuke := "snuke"

	if Ss[0][0] != snuke[0] {
		fmt.Println("No")
		os.Exit(0)
	}

	w_skip = make(map([int]map[int]int))

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			Ss[h][w]
		}
	}

	for {
		i++
		var archives2 [][]int
		for _, wh := range archives {
			h := wh[0]
			w := wh[1]

			if w-1 > 0 {
				if Ss[h][w-1] == snuke[i%5] {
					archives2 = append(archives2, []int{h, w - 1})
				}
			}

			if w+1 < W {
				if Ss[h][w+1] == snuke[i%5] {
					if w+1 == W-1 && h == H-1 {
						fmt.Println("Yes")
						os.Exit(0)
					}

					archives2 = append(archives2, []int{h, w + 1})
				}
			}

			if h-1 > 0 {
				if Ss[h-1][w] == snuke[i%5] {
					archives2 = append(archives2, []int{h - 1, w})
				}
			}

			if h+1 < H {
				if Ss[h+1][w] == snuke[i%5] {
					if w == W-1 && h+1 == H-1 {
						fmt.Println("Yes")
						os.Exit(0)
					}

					archives2 = append(archives2, []int{h + 1, w})
				}
			}
		}

		if len(archives2) == 0 {
			fmt.Println("No")
			os.Exit(0)
		}

		archives = archives2
	}
}

if snukeIndex(s string, snuke string) int {
	for i := 0 i < 4; i++ {
		if s == string(snuke[i]) {
			return i
		}
	}

	return 0
}

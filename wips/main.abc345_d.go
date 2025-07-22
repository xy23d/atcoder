package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N, H, W int
	fmt.Fscan(reader, &N, &H, &W)

	as := make([]int, N)
	bs := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &as[i], &bs[i])
	}

	if stamp(H * W, 0, as, bs, N, H, W, []int{}) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func stamp(z int, index int, as []int, bs []int, N int, H int, W int, is []int) bool {
	for i := index; i < N; i++ {
		if z - as[i] * bs[i] == 0 {
			r := check(H, W, as, bs, append([]int{i}, is...))
			if r {
				return r
			}
		}

		if z > as[i] * bs[i] {
			r := stamp(z - as[i] * bs[i], index + 1, as, bs, N, H, W, append([]int{i}, is...))
			if r {
				return true
			}
		}
	}

	return false
}

func check(H int, W int, as []int, bs []int, is []int) bool {
	c := make(map[int]map[int]bool)

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if c[h][w] {
				continue
			}

			for _, i := range is {
				if (h + as[i]) < H && (w + bs[i]) < W {
					r := true
					for k := h; k < h + as[i]; k++ {
						for j := w; j < w + bs[i]; j++ {
							if c[k][j] {
								r = false
							}
						}
					}

					if r {
						for k := h; k < h + as[i]; k++ {
							for j := w; j < w + bs[i]; j++ {
								c[k][j] = true
							}
						}

						continue
					}
				}

				if (h + bs[i] < H) && (w + as[i] < W) {
					r2 := true
					for k := h; k < h + bs[i]; k++ {
						for j := w; j < w + as[i]; j++ {
							if c[k][j] {
								r2 = false
							}
						}
					}

					if r2 {
						for k := h; k < h + bs[i]; k++ {
							for j := w; j < w + as[i]; j++ {
								c[k][j] = true
							}
						}

						continue
					}
				}

				return false
			}
		}
	}

	return true
}
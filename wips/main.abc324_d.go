package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N int
	var S string
	fmt.Fscan(reader, &N, &S)

	ss := strings.Split(S, "")

	pows := getPows(N)

	fmt.Println(pattern(ss, []int{}, make(map[int]bool), make(map[int]bool), pows))
}

func pattern(ss []string, is []int, a map[int]bool, at map[int]bool, pows map[int]bool) int {
	count := 0

	for i := 0; i < len(ss); i++ {
		
		if a[i] {
			continue
		}

		t2 := make(map[int]bool)
		for k, v := range a {
			t2[k] = v
		}
		t2[i] = true

		t := append(is, i)

		n := calcT(t, ss)

		if len(ss) == len(is) {
			if at[n] {
				continue
			}

			at[n] = true

			if pows[n] {
				count++
			}
		} else {
			count = count + pattern(ss, is, t2, at, pows)
		}
	}

	return count
}

func in(n int, ns []int) bool {
	for i := 0; i < len(ns); i++ {
		if ns[i] == n {
			return true
		}
	}

	return false
}

func calcT(is []int, ss []string) int {
	total := 0

	for i := len(is) - 1; i >= 0 ; i-- {
		index := is[i]

		n, _ := strconv.Atoi(ss[index])

		total += int(math.Pow(10, float64(i))) * n
	}

	return total
}

func getPows(N int) map[int]bool {
	r := make(map[int]bool)

	max := int(math.Pow(10, float64(N)))

	i := 1
	for {
		if i*i > max {
			break
		}

		r[i*i] = true
		i++
	}

	return r
}
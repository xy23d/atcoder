// https://atcoder.jp/contests/adt_hard_20251223_1/tasks/abc320_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var M int
	fmt.Fscan(in, &M)

	var S1, S2, S3 string
	fmt.Fscan(in, &S1, &S2, &S3)

	chars1 := strings.Split(S1, "")
	chars2 := strings.Split(S2, "")
	chars3 := strings.Split(S3, "")

	ns1 := make([][]int, 10)
	ns2 := make([][]int, 10)
	ns3 := make([][]int, 10)
	for n := 0; n < 10; n++ {
		ns1[n] = []int{}
		ns2[n] = []int{}
		ns3[n] = []int{}
	}

	for i := 0; i < M; i++ {
		n, _ := strconv.Atoi(chars1[i])

		ns1[n] = append(ns1[n], i)
		ns1[n] = append(ns1[n], i+M)
		ns1[n] = append(ns1[n], i+M*2)
	}

	for i := 0; i < M; i++ {
		n, _ := strconv.Atoi(chars2[i])

		ns2[n] = append(ns2[n], i)
		ns2[n] = append(ns2[n], i+M)
		ns2[n] = append(ns2[n], i+M*2)
	}

	for i := 0; i < M; i++ {
		n, _ := strconv.Atoi(chars3[i])

		ns3[n] = append(ns3[n], i)
		ns3[n] = append(ns3[n], i+M)
		ns3[n] = append(ns3[n], i+M*2)
	}

	for n := 0; n < 10; n++ {
		sort.Ints(ns1[n])
		sort.Ints(ns2[n])
		sort.Ints(ns3[n])
	}

	ans := -1
	for n := 0; n < 10; n++ {
		if len(ns1[n]) == 0 || len(ns2[n]) == 0 || len(ns3[n]) == 0 {
			continue
		}

		var t int
		t = check(ns1[n], ns2[n], ns3[n])
		if ans == -1 || (t != -1 && ans > t) {
			ans = t
		}

		t = check(ns1[n], ns3[n], ns2[n])
		if ans == -1 || (t != -1 && ans > t) {
			ans = t
		}

		t = check(ns2[n], ns1[n], ns3[n])
		if ans == -1 || (t != -1 && ans > t) {
			ans = t
		}

		t = check(ns2[n], ns3[n], ns1[n])
		if ans == -1 || (t != -1 && ans > t) {
			ans = t
		}

		t = check(ns3[n], ns1[n], ns2[n])
		if ans == -1 || (t != -1 && ans > t) {
			ans = t
		}

		t = check(ns3[n], ns2[n], ns1[n])
		if ans == -1 || (t != -1 && ans > t) {
			ans = t
		}
	}

	fmt.Fprintln(out, ans)
}

func check(ns1, ns2, ns3 []int) int {
	b1 := ns1[0]

	b2 := -1
	for i := 0; i < len(ns2); i++ {
		if b1 != ns2[i] {
			b2 = ns2[i]
			break
		}
	}

	if b2 == -1 {
		return -1
	}

	b3 := -1
	for i := 0; i < len(ns3); i++ {
		if b1 != ns3[i] && b2 != ns3[i] {
			b3 = ns3[i]
			break
		}
	}

	if b3 == -1 {
		return -1
	}

	return max(max(b1, b2), b3)
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var k int64
	fmt.Fscan(in, &k)

	digits := make([]byte, 0, 64)
	for k > 0 {
		if k%2 == 0 {
			digits = append(digits, '0')
		} else {
			digits = append(digits, '2')
		}
		k /= 2
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	fmt.Fprintln(out, string(digits))
}

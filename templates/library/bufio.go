package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var sc = bufio.NewScanner(os.Stdin)

	var N, Q int
	fmt.Fscan(reader, &N, &Q)

	fmt.Fprintln(out, "Hello, world!")
	fmt.Fprint(out, 123, " ", 456, "\n")

	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(out, n)
}

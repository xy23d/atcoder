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

	var N, Q int
	fmt.Fscan(reader, &N, &Q)

	fmt.Fprintln(out, "Hello, world!")
	fmt.Fprint(out, 123, " ", 456, "\n")
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(reader, &N, &Q)
}
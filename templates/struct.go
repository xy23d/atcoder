package main

import (
	"fmt"
)

type Row struct {
	index int
	a     int
	b     int
}

func main() {
	row := Row{1, 2, 3}

	fmt.Println(row)
}

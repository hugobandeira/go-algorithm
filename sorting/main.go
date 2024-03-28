package main

import (
	"fmt"

	"github.com/hugobandeira/go-algorithm/sorting/selection"
)

func main() {
	arr := []int{5, 3, 6, 2, 10}

	fmt.Println(selection.Sort(arr))
}

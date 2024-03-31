package main

import (
	"fmt"

	"github.com/nandanabhishek/filter"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := []int{3, 4, 55, 6, 1, 2}
	filtered := filter.Filter(s1, s2)
	fmt.Printf("Filtered slice %v", filtered)
}

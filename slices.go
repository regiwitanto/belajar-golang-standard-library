package main

import (
	"fmt"
	"slices"
)

func main()  {
	names := []string{"a", "b", "c"}
	values := []int{1, 2, 3}

	// This is a slice of slices
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "a"))
	fmt.Println(slices.Contains(values, 1))
	fmt.Println(slices.Index(names, "b"))
	fmt.Println(slices.Index(values, 2))
}
package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("Regi Witanto", "Regi"))
	fmt.Println(strings.Split("Regi Witanto", " "))
	fmt.Println(strings.ToLower("Regi Witanto"))
	fmt.Println(strings.ToUpper("Regi Witanto"))
	fmt.Println(strings.Trim("         Regi Witanto       ", " "))
	fmt.Println(strings.ReplaceAll("Regi Witanto", "Regi", "Budi"))
}
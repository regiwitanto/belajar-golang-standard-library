package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main()  {
	users := UserSlice{
		{"Tom", 25},
		{"Jerry", 23},
		{"Alice", 27},
	}

	sort.Sort(users)

	fmt.Println(users)
}
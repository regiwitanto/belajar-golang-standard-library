package main

import (
	"container/list"
	"fmt"
)

func main()  {
	var data *list.List = list.New()

	data.PushBack("A")
	data.PushBack("B")
	data.PushBack("C")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
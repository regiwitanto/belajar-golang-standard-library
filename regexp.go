package main

import (
	"fmt"
	"regexp"
)



func main()  {
	var regexp *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regexp.MatchString("eko"))
	fmt.Println(regexp.MatchString("eto"))
	fmt.Println(regexp.MatchString("eDo"))

	fmt.Println(regexp.FindAllString("eko eDo eto ebo", 10))
}
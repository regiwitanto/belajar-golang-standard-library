package main

import (
	"bufio"
	"os"
)

func main()  {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("this is long string\n")
	_, _ = writer.WriteString("with new line\n")
	_ = writer.Flush()
}
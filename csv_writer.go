package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"a", "b", "c"})
	_ = writer.Write([]string{"1", "2", "3"})

	writer.Flush()
}
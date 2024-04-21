package main

import (
	"encoding/csv"
	"strings"
	"io"
	"fmt"
)

func main() {
	csvString := "regi, witanto, aktor\n" +
		"joko, widodo, presiden\n" +
		"jusuf, kalla, wakil presiden"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
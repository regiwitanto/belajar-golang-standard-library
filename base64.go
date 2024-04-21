package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {
	value := "Regi Witanto"

	// Convert string to base64
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("Encoded:", encoded)

	// Convert base64 to string
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Decoded:", string(decoded))
	}
}
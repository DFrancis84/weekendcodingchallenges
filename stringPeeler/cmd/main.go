package main

import (
	"fmt"
	"os"
	"weekendcodingchallenges/stringPeeler/internal/stringpeeler"
)

func main() {
	sP := stringpeeler.New()
	input := os.Args
	input = input[1:]
	fmt.Println("The following words will now be stripped of first and last letter....", input)
	for _, i := range input {
		result, err := sP.StringPeeler(i)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Before StringPeeler: %v, After StringPeeler: %v\n", i, result)
		}
	}

}

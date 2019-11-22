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
	fmt.Println("The following words will now be stripped of first and last letter.... ", input)
	for _, i := range input {
		result := sP.StringPeeler(i)
		fmt.Println(result)
	}

}

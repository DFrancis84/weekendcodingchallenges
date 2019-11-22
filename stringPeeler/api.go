package stringpeeler

import (
	"fmt"
)

type StringPeelerClient struct{}

func New() *StringPeelerClient {
	return &StringPeelerClient{}
}

// func main() {
// 	input := os.Args
// 	input = input[1:]
// 	fmt.Println("The following words will now be stripped of first and last letter.... ", input)
// 	for _, i := range input {
// 		result := StringPeeler(i)
// 		fmt.Println(result)
// 	}
//
// }
//
func StringPeeler(input string) string {
	if len(input) <= 2 {
		output := fmt.Sprintf("word must be larger than 2 characters long. current word '%v' is only %v characters long\n", input, len(input))
		return output
	}
	output := input[:len(input)-1]
	output = output[1:]
	return output
}

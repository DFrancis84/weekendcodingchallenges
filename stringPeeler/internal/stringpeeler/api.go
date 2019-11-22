package stringpeeler

import (
	"fmt"
)

type StringPeelerAPI struct{}

func New() *StringPeelerAPI {
	return &StringPeelerAPI{}
}

func (sp *StringPeelerAPI) StringPeeler(input string) string {
	if len(input) <= 2 {
		output := fmt.Sprintf("word must be larger than 2 characters long. current word '%v' is only %v characters long\n", input, len(input))
		return output
	}
	output := input[:len(input)-1]
	output = output[1:]
	return output
}

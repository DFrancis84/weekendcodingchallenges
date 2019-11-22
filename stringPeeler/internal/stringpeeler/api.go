package stringpeeler

import (
	"errors"
	"fmt"
)

type StringPeelerAPI struct{}

func New() *StringPeelerAPI {
	return &StringPeelerAPI{}
}

func (sp *StringPeelerAPI) StringPeeler(input string) (string, error) {
	if len(input) <= 2 {
		return "", errors.New(fmt.Sprintf("word must be larger than 2 characters long. current word '%v' is only %v characters long", input, len(input)))
	}
	output := input[:len(input)-1]
	output = output[1:]
	return output, nil
}

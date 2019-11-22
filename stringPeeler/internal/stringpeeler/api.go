package stringpeeler

import (
	"errors"
	"fmt"
	"time"
)

type StringPeelerAPI struct{}

func New() *StringPeelerAPI {
	return &StringPeelerAPI{}
}

func (sp *StringPeelerAPI) StringPeeler(input string) (string, error) {
	startTime := time.Now()
	if len(input) <= 2 {
		return "", errors.New(fmt.Sprintf("word must be larger than 2 characters long. current word '%v' is only %v characters long", input, len(input)))
	}
	output := input[:len(input)-1]
	output = output[1:]
	fmt.Printf("Peeler took %v to complete request for %v\n", time.Since(startTime), input)
	return output, nil
}

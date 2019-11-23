package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"weekendcodingchallenges/evenfibonacci/internal/evenfibonacci"
)

func main() {
	startTime := time.Now()
	input := os.Args
	maxNumber, err := strconv.Atoi(input[1])
	if err != nil {
		panic(err)
	}
	eF := evenfibonacci.New()
	list := eF.PopulateFibonacciList(maxNumber)
	sum := eF.CalculateEvenNumbers(list)
	fmt.Printf("The sum of all even numbers in the Fibonacci List capped at %v is %v. Duration of process: %v\n", maxNumber, sum, time.Since(startTime))
}

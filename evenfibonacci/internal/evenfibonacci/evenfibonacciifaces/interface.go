package evenfibonacciifaces

type EvenFibonacciClient interface {
	PopulateFibonacciList(maxNumber int) []int
	CalculateEvenNumbers(list []int) int
}

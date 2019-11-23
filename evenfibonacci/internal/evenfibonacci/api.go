package evenfibonacci

type EvenFibonacciAPI struct{}

func New() *EvenFibonacciAPI {
	return &EvenFibonacciAPI{}
}

func (e *EvenFibonacciAPI) PopulateFibonacciList(maxNumber int) []int {
	fibonacciList := []int{0, 1}
	for i := 2; i < maxNumber; i++ {
		if i == 2 {
			newNumber := fibonacciList[(i-1)] + fibonacciList[(i-1)]
			fibonacciList = append(fibonacciList, newNumber)
			continue
		}
		newNumber := fibonacciList[(i-1)] + fibonacciList[(i-2)]
		if newNumber > maxNumber {
			break
		}
		fibonacciList = append(fibonacciList, newNumber)
	}
	return fibonacciList
}

func (e *EvenFibonacciAPI) CalculateEvenNumbers(list []int) int {
	var sum int
	for _, number := range list {
		if number%2 == 0 {
			sum = sum + number
		} else {
			continue
		}
	}
	return sum
}

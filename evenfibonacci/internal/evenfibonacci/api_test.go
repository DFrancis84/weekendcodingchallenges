package evenfibonacci

import (
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	eF := New()
	want := &EvenFibonacciAPI{}
	if !reflect.DeepEqual(eF, want) {
		t.Errorf("Error creating EvenFibonacciAPI = %v, want = %v", eF, want)
	}
}

func Test_PopulateFibonacciList(t *testing.T) {
	result := []int{0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987}
	type args struct {
		maxNumber int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "PopulateFibonacciList",
			args: args{
				maxNumber: 1000,
			},
			want: result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eF := New()
			got := eF.PopulateFibonacciList(tt.args.maxNumber)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopulateFibonacciList = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_CalculateEvenNumbers(t *testing.T) {
	fibonacciList := []int{0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987}
	result := 798
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CalculateEvenNumbers",
			args: args{
				list: fibonacciList,
			},
			want: result,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eF := New()
			sum := eF.CalculateEvenNumbers(tt.args.list)
			if !reflect.DeepEqual(sum, tt.want) {
				t.Errorf("CalculateEvenNumbers = %v, want = %v", sum, tt.want)
			}
		})
	}

}

package stringpeeler

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	_ = New()
}

func Test_StringPeeler(t *testing.T) {
	type args struct {
		word string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test_StringPeeler",
			args: args{
				word: "alphabet",
			},
			want:    "lphabe",
			wantErr: false,
		},
		{
			name: "Test_StringPeeler_Return_Invalid_Character_Length",
			args: args{
				word: "a",
			},
			want:    "word must be larger than 2 characters long. current word 'a' is only 1 characters long",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := StringPeeler(tt.args.word)
			if reflect.DeepEqual(res, tt.want) {
				fmt.Printf("Want: %v Got: %v\n", tt.want, res)
			}
		})
	}
}

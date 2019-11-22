package stringpeeler

import (
	"fmt"
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
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sP := StringPeelerAPI{}
			fmt.Printf("Word to peel: %v\n", tt.args.word)
			res, err := sP.StringPeeler(tt.args.word)
			if res != tt.want {
				t.Errorf("StringPeeler = %v, want = %v", res, tt.want)
			}
			if err != nil {
				fmt.Println(err)
				if !tt.wantErr {
					t.Errorf("StringPeeler error = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}

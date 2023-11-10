package easy

import (
	"leetcode/problems/easy"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "n less than 1",
			args: args{n: 0},
			want: []string{},
		},
		{
			name: "n equals 1",
			args: args{n: 1},
			want: []string{"1"},
		},
		{
			name: "n equals 3",
			args: args{n: 3},
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "n equals 5",
			args: args{n: 5},
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "n equals 15",
			args: args{n: 15},
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.FizzBuzz(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

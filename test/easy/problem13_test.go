package easy

import (
	"leetcode/problems/easy"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Case I", args: args{s: "I"}, want: 1},
		{name: "Case IV", args: args{s: "IV"}, want: 4},
		{name: "Case IX", args: args{s: "IX"}, want: 9},
		{name: "Case LVIII", args: args{s: "LVIII"}, want: 58},
		{name: "Case MCMXCIV", args: args{s: "MCMXCIV"}, want: 1994},
		{name: "Case XLVII", args: args{s: "XLVII"}, want: 47},
		{name: "Case CM", args: args{s: "CM"}, want: 900},
		{name: "Case MMMCMXCIX", args: args{s: "MMMCMXCIX"}, want: 3999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := easy.RomanToInt(tt.args.s); got != tt.want {
				t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

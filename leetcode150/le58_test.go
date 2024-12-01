package leetcode150

import (
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first", args{"hello world"}, 5},
		{"second", args{"  fly me   to   the moon  "}, 4},
		{"third", args{"a "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLastWord2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first", args{"hello world"}, 5},
		{"second", args{"  fly me   to   the moon  "}, 4},
		{"third", args{"a "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord2() = %v, want %v", got, tt.want)
			}
		})
	}
}

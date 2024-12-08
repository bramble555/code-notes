package leetcode150

import (
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	type res struct {
		nums []int
		ret  int
	}
	tests := []struct {
		name string
		args args
		want res
	}{
		// TODO: Add test cases.
		{"first", args{[]int{1, 2, 2, 5}, 2}, res{[]int{1, 5, 2, 2}, 2}},
		{"second", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, res{[]int{0, 1, 3, 0, 4, 2, 2, 2}, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("\ntt.args.nums:", tt.args.nums)
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want.ret {
				t.Log("\ntt.args.nums:", tt.args.nums)
				t.Errorf("\nremoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

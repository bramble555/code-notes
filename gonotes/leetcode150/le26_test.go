package leetcode150

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
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
		{"first", args{[]int{1, 2, 2, 5}}, res{[]int{1, 2, 5, 2}, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("\ntt.args.nums at frist:", tt.args.nums)
			got := removeDuplicates(tt.args.nums)
			t.Log("\ntt.args.nums in the end:", tt.args.nums)
			if got != tt.want.ret {

				t.Errorf("\nremoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode150

import "testing"

func Test_removeDuplicates2(t *testing.T) {
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
		{"first", args{[]int{1, 1, 1, 2, 2, 3}}, res{[]int{1, 1, 2, 2, 3}, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("\ntt.args.nums at first:", tt.args.nums)
			got := removeDuplicates2(tt.args.nums)
			t.Log("\ntt.args.nums int the end:", tt.args.nums)
			if got != tt.want.ret {
				t.Errorf("\nremoveDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}

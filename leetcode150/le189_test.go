package leetcode150

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"first", args{[]int{-1, -100, 3, 99}, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("tt.args.nums at first:%v", tt.args.nums)
			rotate(tt.args.nums, tt.args.k)
			t.Logf("tt.args.nums in the end:%v", tt.args.nums)
		})
	}
}

package leetcode150

import (
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first", args{[]int{3, 2, 3}}, 3},
		{"second", args{[]int{3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("tt.args.nums at first:%v", tt.args.nums)
			got := majorityElement(tt.args.nums)
			t.Logf("tt.args.nums in the end:%v", tt.args.nums)
			if got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode150

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"first", args{[]int{2, 3, 1, 1, 4}}, true},
		{"first2", args{[]int{2, 3, 1, 0, 4}}, true},
		{"second", args{[]int{2, 0, 0}}, true},
		{"second2", args{[]int{1, 0, 0}}, false},
		{"third", args{[]int{2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode150

import "testing"

func Test_maxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// 4 + 3 = 7
		{"first", args{[]int{7, 1, 5, 3, 6, 4}}, 7},
		{"first", args{[]int{7}}, 0},
		{"first", args{[]int{7, 1}}, 0},
		{"sceod", args{[]int{1, 2, 3, 4, 5}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode150

import "testing"

func Test_rotate2(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"first",
			args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate2(tt.args.matrix)
		})
	}
}

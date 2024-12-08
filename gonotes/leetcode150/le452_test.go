package leetcode150

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {"first", args{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}}, 2},
		// {"first", args{[][]int{{2, 6}, {3, 8}, {7, 10}}}, 2},
		// {"first", args{[][]int{{2, 8}, {2, 9}, {3, 6}}}, 1},
		{"forth", args{[][]int{{1, 2}, {3, 4}, {0, 6}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}

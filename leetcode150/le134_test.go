package leetcode150

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {"second", args{[]int{2, 3, 4}, []int{3, 4, 3}}, -1},
		// {"first", args{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}}, 3},
		// {"third", args{[]int{4, 5, 2, 6, 5, 3}, []int{3, 2, 7, 3, 2, 9}}, -1},
		{"forth", args{[]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}}, 4},
		{"forth", args{[]int{5, 1, 2, 4, 4}, []int{4, 4, 1, 5, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

package leetcode150

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"first", args{[]int{3, 0, 6, 1, 5}}, 3},
		{"second", args{[]int{1, 2, 3}}, 2},
		{"third", args{[]int{0}}, 0},
		{"third", args{[]int{0, 1}}, 1},
		{"third", args{[]int{11, 15}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

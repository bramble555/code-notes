package leetcode150

import (
	"reflect"
	"testing"
)

func Test_merge2(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"first", args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge2(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge2() = %v, want %v", got, tt.want)
			}
		})
	}
}

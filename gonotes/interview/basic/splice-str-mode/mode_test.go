package splicestrmode

import "testing"

func Test_mode(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mode()
		})
	}
}

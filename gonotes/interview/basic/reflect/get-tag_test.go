package reflect

import "testing"

func Test_get(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"first"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get()
		})
	}
}

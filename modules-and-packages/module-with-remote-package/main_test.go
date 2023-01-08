package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test_main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

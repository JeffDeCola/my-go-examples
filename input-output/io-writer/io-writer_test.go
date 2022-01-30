package main

import (
	"bytes"
	"testing"
)

func Test_writeFromBuffer(t *testing.T) {
	tests := []struct {
		name    string
		wantW   string
		wantErr bool
	}{
		{
			name:    "Test Write",
			wantW:   "This data is being put into a byte reader\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := writeFromBuffer(w); (err != nil) != tt.wantErr {
				t.Errorf("writeFromBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("writeFromBuffer() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Main",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

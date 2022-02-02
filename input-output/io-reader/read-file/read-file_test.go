package main

import (
	"io"
	"os"
	"testing"
)

func openFile(filename string) *os.File {
	r, _ := os.Open(filename)
	return r
}

func Test_readFile(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test read file - error",
			args: args{
				r: openFile("nothing.txt"),
			},
			wantErr: true,
		},
		{
			name: "Test read file",
			args: args{
				r: openFile("test.txt"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := readFile(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

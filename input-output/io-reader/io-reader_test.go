package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func openFile(filename string) *os.File {
	r, _ := os.Open(filename)
	return r
}

func Test_readToBufferandPrint(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test string",
			args: args{
				r: strings.NewReader("Hi everyone, this is a string"),
			},
			wantErr: false,
		},
		{
			name: "Test bad io.Reader with a bad open file",
			args: args{
				r: openFile("nothing.txt"),
			},
			wantErr: true,
		},
		{
			name: "Test string stop",
			args: args{
				r: strings.NewReader("stop\n"),
			},
			wantErr: false,
		},
		{
			name: "Test buffer",
			args: args{
				r: strings.NewReader("Hi everyone, this is a string"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := readToBufferandPrint(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("readToBufferandPrint() error = %v, wantErr %v", err, tt.wantErr)
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

package main

import (
	"io"
	"strings"
	"testing"
)

func Test_readUsingReadMethod(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test using read method",
			args: args{
				r: strings.NewReader("Hi"),
			},
			want:    "Hi",
			wantErr: false,
		},
		{
			name: "Test using read method - Larger buffer size",
			args: args{
				r: strings.NewReader("Hi this is a very large buffer\n"),
			},
			want:    "Hi this is",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readUsingReadMethod(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readUsingReadMethod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readUsingReadMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readUsingFscan(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test using Fscan method",
			args: args{
				r: strings.NewReader("Hi"),
			},
			want:    "Hi",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readUsingFscan(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readUsingFscan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readUsingFscan() = %v, want %v", got, tt.want)
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

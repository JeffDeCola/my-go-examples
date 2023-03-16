package main

import (
	"io"
	"testing"
)

func Test_setLogLevel(t *testing.T) {
	type args struct {
		logLevel string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setLogLevel(tt.args.logLevel); (err != nil) != tt.wantErr {
				t.Errorf("setLogLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getNumbers(t *testing.T) {
	type args struct {
		r1 io.Reader
		r2 io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getNumbers(tt.args.r1, tt.args.r2)
			if (err != nil) != tt.wantErr {
				t.Errorf("getNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getNumbers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getNumbers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getUserInput(t *testing.T) {
	type args struct {
		r       io.Reader
		askUser string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUserInput(tt.args.r, tt.args.askUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getUserInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertStringToInt(t *testing.T) {
	type args struct {
		nString string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertStringToInt(tt.args.nString)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertStringToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertStringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSum(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
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

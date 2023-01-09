package main

import "testing"

func Test_sendData(t *testing.T) {
	type args struct {
		chData       chan string
		chUnBuffered chan string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendData(tt.args.chData, tt.args.chUnBuffered)
		})
	}
}

func Test_rcvData(t *testing.T) {
	type args struct {
		chUnBuffered chan string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rcvData(tt.args.chUnBuffered)
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

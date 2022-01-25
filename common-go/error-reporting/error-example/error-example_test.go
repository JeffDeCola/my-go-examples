// my-go-examples error-example.go

package main

import (
	"io"
	"strings"
	"testing"
)

func Test_askQuestion(t *testing.T) {
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
			name: "Test Get User Input Error",
			args: args{
				r: strings.NewReader("\n"),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test Get User Input",
			args: args{
				r: strings.NewReader("jeff"),
			},
			want:    "jeff",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := askQuestion(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("askQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("askQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkAnswer(t *testing.T) {
	type args struct {
		answer string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Integer - Correct Number",
			args: args{
				answer: "4",
			},
			wantErr: false,
		},
		{
			name: "Integer - Incorrect Number",
			args: args{
				answer: "5",
			},
			wantErr: true,
		},
		{
			name: "Not and Integer",
			args: args{
				answer: "a",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkAnswer(tt.args.answer); (err != nil) != tt.wantErr {
				t.Errorf("checkAnswer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkNumber(t *testing.T) {
	type args struct {
		answer int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 4",
			args: args{
				answer: 4,
			},
			wantErr: false,
		},
		{
			name: "Test 5 - Error",
			args: args{
				answer: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkNumber(tt.args.answer); (err != nil) != tt.wantErr {
				t.Errorf("checkNumber() error = %v, wantErr %v", err, tt.wantErr)
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

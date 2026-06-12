// my-go-examples error-example

package main

import "testing"

func Test_firstLevel(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test for empty filename", args{""}, true},
		{"test for filename", args{"jeff"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := firstLevel(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("firstLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_secondLevel(t *testing.T) {
	type args struct {
		filename string
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
			if err := secondLevel(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("secondLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_thirdLevel(t *testing.T) {
	type args struct {
		filename string
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
			if err := thirdLevel(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("thirdLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Main test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

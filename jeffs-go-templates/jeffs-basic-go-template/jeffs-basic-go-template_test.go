package main

import "testing"

func Test_setLogLevel(t *testing.T) {
	type args struct {
		logLevel string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Trace",
			args: args{
				logLevel: "trace",
			},
			wantErr: false,
		},
		{
			name: "Test Info",
			args: args{
				logLevel: "info",
			},
			wantErr: false,
		},
		{
			name: "Test Error",
			args: args{
				logLevel: "error",
			},
			wantErr: false,
		},
		{
			name: "Test Bad Log Level",
			args: args{
				logLevel: "whatever",
			},
			wantErr: true,
		},
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
	tests := []struct {
		name    string
		want    int
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getNumbers()
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
		ask string
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
			got, err := getUserInput(tt.args.ask)
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
		{
			name: "Test using a number",
			args: args{
				nString: "2",
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "Test using a string",
			args: args{
				nString: "two",
			},
			want:    -1,
			wantErr: true,
		},
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
		{
			name: "Test 1+1",
			args: args{
				n1: 1,
				n2: 1,
			},
			want: 2,
		},
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

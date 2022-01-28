package main

import "testing"

func Test_areaRectangle(t *testing.T) {
	type args struct {
		w float64
		h float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Area Rectangle",
			args: args{
				w: 2.4,
				h: 34.4,
			},
			want: 82.55999999999999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areaRectangle(tt.args.w, tt.args.h); got != tt.want {
				t.Errorf("areaRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_areaCircle(t *testing.T) {
	type args struct {
		r float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Area Circle",
			args: args{
				r: 2.3,
			},
			want: 16.619025137490002,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areaCircle(tt.args.r); got != tt.want {
				t.Errorf("areaCircle() = %v, want %v", got, tt.want)
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

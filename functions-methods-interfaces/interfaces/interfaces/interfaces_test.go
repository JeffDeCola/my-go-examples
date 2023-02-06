package main

import (
	"testing"
)

func TestRectangle_area(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Area Rectangle",
			fields: fields{
				width:  2.4,
				height: 34.4,
			},
			want: 82.55999999999999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rectangle{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.area(); got != tt.want {
				t.Errorf("Rectangle.theArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Area Circle",
			fields: fields{
				radius: 2.3,
			},
			want: 16.619025137490002,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := circle{
				radius: tt.fields.radius,
			}
			if got := c.area(); got != tt.want {
				t.Errorf("Circle.theArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getArea(t *testing.T) {
	type args struct {
		g geometry
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Area Interface - Rectangle",
			args: args{
				g: rectangle{2.4, 34.4},
			},
			want: 82.55999999999999,
		},
		{
			name: "Test Area Interface - Circle",
			args: args{
				g: circle{2.3},
			},
			want: 16.619025137490002,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getArea(tt.args.g); got != tt.want {
				t.Errorf("area() = %v, want %v", got, tt.want)
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

package main

import "testing"

func TestRectangle_areaRectangle(t *testing.T) {
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
			r := Rectangle{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.areaRectangle(); got != tt.want {
				t.Errorf("Rectangle.areaRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_areaCircle(t *testing.T) {
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
			r := Circle{
				radius: tt.fields.radius,
			}
			if got := r.areaCircle(); got != tt.want {
				t.Errorf("Circle.areaCircle() = %v, want %v", got, tt.want)
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

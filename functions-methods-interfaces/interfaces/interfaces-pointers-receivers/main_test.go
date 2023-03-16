package main

import (
	"testing"
)

func TestRectangle_area(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	type args struct {
		a *float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Area Rectangle",
			fields: fields{
				width:  2.4,
				height: 34.4,
			},
			args: args{
				// want: 82.55999999999999,
				a: new(float64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rectangle{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			r.area(tt.args.a)
		})
	}
}

func TestCircle_area(t *testing.T) {
	type fields struct {
		radius float64
	}
	type args struct {
		a *float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Area Circle",
			fields: fields{
				radius: 2.3,
			},
			args: args{
				// want: 16.619025137490002,
				a: new(float64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := circle{
				radius: tt.fields.radius,
			}
			c.area(tt.args.a)
		})
	}
}

func TestRectangle_size(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	type args struct {
		f float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Size Rectangle",
			fields: fields{
				width:  2.4,
				height: 34.4,
				// want: 330.24,
			},
			args: args{
				f: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rectangle{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			r.size(tt.args.f)
		})
	}
}

func TestCircle_size(t *testing.T) {
	type fields struct {
		radius float64
	}
	type args struct {
		f float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Size Circle",
			fields: fields{
				radius: 2.3,
				// want: 66.48,
			},
			args: args{
				f: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &circle{
				radius: tt.fields.radius,
			}
			c.size(tt.args.f)
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

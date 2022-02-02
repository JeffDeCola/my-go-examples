package geometry

import (
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		Width  float64
		Height float64
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
				Width:  2.4,
				Height: 34.4,
			},
			// want: 82.55999999999999,
			args: args{
				a: new(float64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			r.Area(tt.args.a)
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	type fields struct {
		Width  float64
		Height float64
	}
	type args struct {
		p *float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Perimeter Rectangle",
			fields: fields{
				Width:  2.4,
				Height: 34.4,
			},
			// want: 73.6,
			args: args{
				p: new(float64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			r.Perimeter(tt.args.p)
		})
	}
}

func TestCircle_Area(t *testing.T) {
	type fields struct {
		Radius float64
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
				Radius: 2.3,
			},
			// want: 16.619025137490002,
			args: args{
				a: new(float64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			c.Area(tt.args.a)
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		Radius float64
	}
	type args struct {
		p *float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Perimeter Circle",
			fields: fields{
				Radius: 2.3,
			},
			// want: 14.451326206513047,
			args: args{
				p: new(float64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			c.Perimeter(tt.args.p)
		})
	}
}

func TestTriangle_Area(t *testing.T) {
	type fields struct {
		A float64
		B float64
		C float64
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
			name: "Test Area Triangle",
			fields: fields{
				A: 2,
				B: 3.3,
				C: 4,
			},
			// want: 3.2883116868691165,
			args: args{
				a: new(float64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Triangle{
				A: tt.fields.A,
				B: tt.fields.B,
				C: tt.fields.C,
			}
			tr.Area(tt.args.a)
		})
	}
}

func TestTriangle_Perimeter(t *testing.T) {
	type fields struct {
		A float64
		B float64
		C float64
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
			name: "Test Perimeter Triangle",
			fields: fields{
				A: 2,
				B: 3.3,
				C: 4,
			},
			// want: 9.3,
			args: args{
				a: new(float64),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Triangle{
				A: tt.fields.A,
				B: tt.fields.B,
				C: tt.fields.C,
			}
			tr.Perimeter(tt.args.a)
		})
	}
}

func TestRectangle_Size(t *testing.T) {
	type fields struct {
		Width  float64
		Height float64
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
				Width:  2.4,
				Height: 34.4,
			},
			args: args{
				f: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			r.Size(tt.args.f)
		})
	}
}

func TestCircle_Size(t *testing.T) {
	type fields struct {
		Radius float64
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
				Radius: 2.3,
			},
			args: args{
				f: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				Radius: tt.fields.Radius,
			}
			c.Size(tt.args.f)
		})
	}
}

func TestTriangle_Size(t *testing.T) {
	type fields struct {
		A float64
		B float64
		C float64
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
			name: "Test Size Triangle",
			fields: fields{
				A: 2,
				B: 3.3,
				C: 4,
			},
			args: args{
				f: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Triangle{
				A: tt.fields.A,
				B: tt.fields.B,
				C: tt.fields.C,
			}
			tr.Size(tt.args.f)
		})
	}
}

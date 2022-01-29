package shapes

import "testing"

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Area Rectangle",
			fields: fields{
				Width:  2.4,
				Height: 34.4,
			},
			want: 82.55999999999999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := r.Area(); got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Perimeter Rectangle",
			fields: fields{
				Width:  2.4,
				Height: 34.4,
			},
			want: 73.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := r.Perimeter(); got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	type fields struct {
		Radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Area Circle",
			fields: fields{
				Radius: 2.3,
			},
			want: 16.619025137490002,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			if got := c.Area(); got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		Radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Perimeter Circle",
			fields: fields{
				Radius: 2.3,
			},
			want: 14.451326206513047,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				Radius: tt.fields.Radius,
			}
			if got := c.Perimeter(); got != tt.want {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangle_Area(t *testing.T) {
	type fields struct {
		A float64
		B float64
		C float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Area Triangle",
			fields: fields{
				A: 2,
				B: 3.3,
				C: 4,
			},
			want: 3.2883116868691165,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Triangle{
				A: tt.fields.A,
				B: tt.fields.B,
				C: tt.fields.C,
			}
			if got := tr.Area(); got != tt.want {
				t.Errorf("Triangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangle_Perimeter(t *testing.T) {
	type fields struct {
		A float64
		B float64
		C float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test Perimeter Triangle",
			fields: fields{
				A: 2,
				B: 3.3,
				C: 4,
			},
			want: 9.3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Triangle{
				A: tt.fields.A,
				B: tt.fields.B,
				C: tt.fields.C,
			}
			if got := tr.Perimeter(); got != tt.want {
				t.Errorf("Triangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArea(t *testing.T) {
	type args struct {
		g Geometry
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Area Interface - Rectangle",
			args: args{
				g: Rectangle{2.4, 34.4},
			},
			want: 82.55999999999999,
		},
		{
			name: "Test Area Interface - Circle",
			args: args{
				g: Circle{2.3},
			},
			want: 16.619025137490002,
		},
		{
			name: "Test Area Interface - Triangle",
			args: args{
				g: Triangle{2, 3.3, 4},
			},
			want: 3.2883116868691165,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetArea(tt.args.g); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPerimeter(t *testing.T) {
	type args struct {
		g Geometry
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Perimeter Interface - Rectangle",
			args: args{
				g: Rectangle{2.4, 34.4},
			},
			want: 73.6,
		},
		{
			name: "Test Perimeter Interface - Circle",
			args: args{
				g: Circle{2.3},
			},
			want: 14.451326206513047,
		},
		{
			name: "Test Perimeter Interface - Triangle",
			args: args{
				g: Triangle{2, 3.3, 4},
			},
			want: 9.3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPerimeter(tt.args.g); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

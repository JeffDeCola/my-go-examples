// my-go-examples gotests.go

package rectangle

import "testing"

func TestRectangleAreaFunction(t *testing.T) {
	type args struct {
		r Rectangle
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add some test cases.
		{"testFunction1",
			args{
				r: Rectangle{
					Width:  6.0,
					Height: 3.0,
				},
			},
			18.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RectangleAreaFunction(tt.args.r); got != tt.want {
				t.Errorf("RectangleAreaFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_RectangleAreaMethod(t *testing.T) {
	type fields struct {
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{"testMethod1",
			fields{
				Width:  6.0,
				Height: 3.0,
			},
			18.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := r.RectangleAreaMethod(); got != tt.want {
				t.Errorf("Rectangle.RectangleAreaMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

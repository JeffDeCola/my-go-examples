package creatures

import "testing"

func TestWarewolf_Kind(t *testing.T) {
	type fields struct {
		TimeofDay string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Test1",
			fields{
				TimeofDay: "day",
			},
			"human",
		},
		{
			"Test2",
			fields{
				TimeofDay: "night",
			},
			"warewolf",
		},
	}
	for _, tt := range tests {
		w := Warewolf{
			TimeofDay: tt.fields.TimeofDay,
		}
		if got := w.Kind(); got != tt.want {
			t.Errorf("%q. Warewolf.Kind() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestWarewolf_Fly(t *testing.T) {
	type fields struct {
		TimeofDay string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Test1",
			fields{
				TimeofDay: "day",
			},
			false,
		},
		{
			"Test2",
			fields{
				TimeofDay: "night",
			},
			false,
		},
	}
	for _, tt := range tests {
		w := Warewolf{
			TimeofDay: tt.fields.TimeofDay,
		}
		if got := w.Fly(); got != tt.want {
			t.Errorf("%q. Warewolf.Fly() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestWarewolf_Sound(t *testing.T) {
	type fields struct {
		TimeofDay string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Warewolf is human during day",
			fields{
				TimeofDay: "day",
			},
			"hello",
		},
		{
			"Warewolf is amonster at night",
			fields{
				TimeofDay: "night",
			},
			"howl",
		},
	}
	for _, tt := range tests {
		w := Warewolf{
			TimeofDay: tt.fields.TimeofDay,
		}
		if got := w.Sound(); got != tt.want {
			t.Errorf("%q. Warewolf.Sound() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestVampire_Kind(t *testing.T) {
	type fields struct {
		Age int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Over 100 he's a vampire",
			fields{
				Age: 134,
			},
			"Vampire",
		},
		{
			"Under 100 he's a human",
			fields{
				Age: 34,
			},
			"human",
		},
	}
	for _, tt := range tests {
		z := Vampire{
			Age: tt.fields.Age,
		}
		if got := z.Kind(); got != tt.want {
			t.Errorf("%q. Vampire.Kind() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestVampire_Fly(t *testing.T) {
	type fields struct {
		Age int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Over 100 he can fly",
			fields{
				Age: 134,
			},
			true,
		},
		{
			"Under 100 he can fly",
			fields{
				Age: 34,
			},
			true,
		},
	}
	for _, tt := range tests {
		z := Vampire{
			Age: tt.fields.Age,
		}
		if got := z.Fly(); got != tt.want {
			t.Errorf("%q. Vampire.Fly() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestVampire_Sound(t *testing.T) {
	type fields struct {
		Age int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"Over 100",
			fields{
				Age: 134,
			},
			"I want to drink your blood",
		},
		{
			"Under 100",
			fields{
				Age: 34,
			},
			"I want to drink your blood",
		},
	}
	for _, tt := range tests {
		z := Vampire{
			Age: tt.fields.Age,
		}
		if got := z.Sound(); got != tt.want {
			t.Errorf("%q. Vampire.Sound() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

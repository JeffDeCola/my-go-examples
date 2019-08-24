package creatures

import "testing"

func TestWerewolf_Kind(t *testing.T) {
	type fields struct {
		TimeofDay string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
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
			"werewolf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Werewolf{
				TimeofDay: tt.fields.TimeofDay,
			}
			if got := w.Kind(); got != tt.want {
				t.Errorf("Werewolf.Kind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWerewolf_Fly(t *testing.T) {
	type fields struct {
		TimeofDay string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
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
		t.Run(tt.name, func(t *testing.T) {
			w := Werewolf{
				TimeofDay: tt.fields.TimeofDay,
			}
			if got := w.Fly(); got != tt.want {
				t.Errorf("Werewolf.Fly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWerewolf_Sound(t *testing.T) {
	type fields struct {
		TimeofDay string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			"Test1",
			fields{
				TimeofDay: "day",
			},
			"hello",
		},
		{
			"Test2",
			fields{
				TimeofDay: "night",
			},
			"howl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Werewolf{
				TimeofDay: tt.fields.TimeofDay,
			}
			if got := w.Sound(); got != tt.want {
				t.Errorf("Werewolf.Sound() = %v, want %v", got, tt.want)
			}
		})
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
		// TODO: Add test cases.
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
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := Vampire{
				Age: tt.fields.Age,
			}
			if got := z.Kind(); got != tt.want {
				t.Errorf("Vampire.Kind() = %v, want %v", got, tt.want)
			}
		})
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
		// TODO: Add test cases.
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
		t.Run(tt.name, func(t *testing.T) {
			z := Vampire{
				Age: tt.fields.Age,
			}
			if got := z.Fly(); got != tt.want {
				t.Errorf("Vampire.Fly() = %v, want %v", got, tt.want)
			}
		})
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
		// TODO: Add test cases.
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
		t.Run(tt.name, func(t *testing.T) {
			z := Vampire{
				Age: tt.fields.Age,
			}
			if got := z.Sound(); got != tt.want {
				t.Errorf("Vampire.Sound() = %v, want %v", got, tt.want)
			}
		})
	}
}

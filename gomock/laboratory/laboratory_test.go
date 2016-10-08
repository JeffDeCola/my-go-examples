package laboratory

import "testing"

func TestCreatureSound(t *testing.T) {
	type args struct {
		scarysound Creatures
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := CreatureSound(tt.args.scarysound); got != tt.want {
			t.Errorf("%q. CreatureSound() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCreatureFly(t *testing.T) {
	type args struct {
		canfly Creatures
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := CreatureFly(tt.args.canfly); got != tt.want {
			t.Errorf("%q. CreatureFly() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

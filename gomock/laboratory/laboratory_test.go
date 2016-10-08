package laboratory

import "testing"

func TestSqeezeDuck(t *testing.T) {
	type args struct {
		ducktosqeeze Duck
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := SqeezeDuck(tt.args.ducktosqeeze); got != tt.want {
			t.Errorf("%q. SqeezeDuck() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestKillLivingThing(t *testing.T) {
	type args struct {
		thingtokill LivingThing
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := KillLivingThing(tt.args.thingtokill); got != tt.want {
			t.Errorf("%q. KillLivingThing() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

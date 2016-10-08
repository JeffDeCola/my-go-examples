package laboratory

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGreet(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	var mockcreature = NewMockCreatures(ctrl)
	mockcreature.EXPECT().Sound().Times(1).Return("poo")
	type args struct {
		c Creatures
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test1",
			args{
				c: mockcreature,
			},
			"poo",
		},
	}
	for _, tt := range tests {
		if got := Greet(tt.args.c); got != tt.want {
			t.Errorf("%q. Greet() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFlyAway(t *testing.T) {
	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	var mockcreature = NewMockCreatures(ctrl)
	gomock.InOrder(
		mockcreature.EXPECT().Fly().Times(1).Return(true),
		mockcreature.EXPECT().Fly().Times(1).Return(false),
	)
	mockcreature.EXPECT().Kind().AnyTimes().Return("Pig")
	type args struct {
		canfly Creatures
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Creature Can Fly",
			args{
				canfly: mockcreature,
			},
			"Pig's can Fly - flap flap flap",
		},
		{
			"Creature Can't Fly",
			args{
				canfly: mockcreature,
			},
			"Pig's can't Fly",
		},
	}
	for _, tt := range tests {
		if got := FlyAway(tt.args.canfly); got != tt.want {
			t.Errorf("%q. FlyAway() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

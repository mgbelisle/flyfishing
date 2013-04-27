package flyfishing

import (
	"math"
	"math/rand"
)

type Lake struct {
	MaxLocation Location
	Fishes []Fish
}

func (l Lake) CastFly(fly Fly, loc Location) Fish {
	for _, fish := range l.Fishes {
		distance := math.Sqrt(math.Pow(loc.X, 2) + math.Pow(loc.Y, 2))
		if fish.LureWith(fly, distance) {
			return fish
		}
	}
	return nil
}
func (l *Lake) AddFish() {
	randX := rand.Float64() * l.MaxLocation.X
	randY := rand.Float64() * l.MaxLocation.Y
	l.Fishes = append(l.Fishes, Rainbow{trout{Location{randX, randY}}})
}

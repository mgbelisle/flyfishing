package flyfishing

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Lake struct {
	Length, Width float64
	Fishes []Fish
}
func (l Lake) CastInto(fly Fly, loc Location) Fish {
	for _, fish := range l.Fishes {
		distance := math.Sqrt(math.Pow(loc.X, 2) + math.Pow(loc.Y, 2))
		if fish.LureWith(fly, distance) {
			return fish
		}
	}
	return nil
}
func (l Lake) RandLoc() Location {
	x := rand.Float64() * l.Length
	y := rand.Float64() * l.Width
	return Location{x, y}
}
func (l Lake) NewFish() Fish {
	r := rand.Float64()
	if r < 0.3 {
		return Rainbow{Trout{l.RandLoc()}}
	}
	return Cutthroat{Trout{l.RandLoc()}}
}
func NewLake() Lake {
	length := 200 + rand.Float64() * 800
	width := 200 + rand.Float64() * 800
	lake := Lake{Length: length, Width: width}
	numFishes := rand.Float64() * length * width / 10
	for i := 0; i < int(numFishes) ; i++ {
		lake.Fishes = append(lake.Fishes, lake.NewFish())
	}
	return lake
}

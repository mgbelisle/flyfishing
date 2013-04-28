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
	length, width float64
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
	x := rand.Float64() * l.length
	y := rand.Float64() * l.width
	return Location{x, y}
}
func (l Lake) NewFish() Fish {
	r := rand.Float64()
	if r < 0.3 {
		return Rainbow{Trout{l.RandLoc()}}
	}
	return Cutthroat{Trout{l.RandLoc()}}
}
func NewLake(length, width float64) Lake {
	// length := rand.Float64() * 1000
	// width := rand.Float64() * 1000
	lake := Lake{length: length, width: width}
	numFishes := rand.Float64() * lake.length * lake.width / 10
	for i := 0; i < int(numFishes) ; i++ {
		lake.Fishes = append(lake.Fishes, lake.NewFish())
	}
	return lake
}

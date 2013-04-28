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
	fishes []Fish
}
func (l Lake) CastInto(fly Fly, loc Location) Fish {
	for _, fish := range l.fishes {
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
func (l Lake) NewFish(loc Location) Fish {
	r := rand.Float64()
	if r < 0.3 {
		return Rainbow{Trout{loc}}
	}
	return Cutthroat{Trout{loc}}
}
func NewLake() Lake {
	lake := Lake{Length: 500, Width: 300}
	sweetSpot := lake.RandLoc()
	// Adds 100 fish to the lake randomly placed around the sweet
	// spot.
	for i := 0; i < 100 ; i++ {
		loc1 := lake.RandLoc()
		x := sweetSpot.X - (sweetSpot.X - loc1.X) / 5
		y := sweetSpot.Y - (sweetSpot.Y - loc1.Y) / 5
		lake.fishes = append(lake.fishes, lake.NewFish(Location{x, y}))
	}
	return lake
}

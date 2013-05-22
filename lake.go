package flyfishing

import (
	"math/rand"
	"time"
)

// init() runs once a package is imported.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Lake is like a class.
type Lake struct {
	Length, Width float64
	fishes        []Fish
}

// Structs support methods like this.
func (l Lake) CastInto(fly Fly, loc Location) Fish {
	time.Sleep(time.Millisecond * 100)
	// range returns index, element for every element in an
	// iterable.  The underscore ignores a variable if you don't
	// need it.
	for _, fish := range l.fishes {
		if fish.lureWith(fly, loc) {
			return fish
		}
	}
	return nil
}

// Gets a random location on the lake
func (l Lake) RandLoc() Location {
	x := rand.Float64() * l.Length
	y := rand.Float64() * l.Width
	return Location{x, y}
}

// Adds a new fish to the lake
func (l Lake) newFish(loc Location) Fish {
	r := rand.Float64()
	if r < 0.5 {
		return Rainbow{Trout{loc}}
	}
	return Cutthroat{Trout{loc}}
}

// Instantiates a new lake.  The convention is to have NewClassName()
// but you can name it whatever you want and give it whatever
// parameters you want.
func NewLake() Lake {
	lake := Lake{Length: 500, Width: 300}
	sweetSpot := lake.RandLoc()
	// Adds 1000 fish to the lake placed randomly around a sweet
	// spot.
	for i := 0; i < 1000; i++ {
		loc := lake.RandLoc()
		r := rand.Float64()
		loc.X = sweetSpot.X + (loc.X-sweetSpot.X)*r
		loc.Y = sweetSpot.Y + (loc.Y-sweetSpot.Y)*r
		lake.fishes = append(lake.fishes, lake.newFish(loc))
	}
	return lake
}

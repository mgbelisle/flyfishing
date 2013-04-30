package flyfishing

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Lake is like a class
type Lake struct {
	Length, Width float64
	fishes []Fish
}
func (l Lake) CastInto(fly Fly, loc Location) Fish {
	time.Sleep(time.Millisecond * 100)
	for _, fish := range l.fishes {
		distance := math.Sqrt(math.Pow(loc.X, 2) + math.Pow(loc.Y, 2))
		if fish.lureWith(fly, distance) {
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
func (l Lake) newFish(loc Location) Fish {
	r := rand.Float64()
	if r < 0.3 {
		return Rainbow{Trout{loc}}
	}
	return Cutthroat{Trout{loc}}
}

func NewLake() Lake {
	lake := Lake{Length: 500, Width: 300}
	sweetSpot := lake.RandLoc()
	// Adds 1000 fish to the lake placed randomly around a sweet
	// spot.
	locations := []Location{}
	for i := 0; i < 1000 ; i++ {
		loc := lake.RandLoc()
		r := rand.Float64()
		loc.X = sweetSpot.X + (loc.X - sweetSpot.X) * r
		loc.Y = sweetSpot.Y + (loc.Y - sweetSpot.Y) * r
		locations = append(locations, loc)
		lake.fishes = append(lake.fishes, lake.newFish(loc))
	}
	lake.LocationsToSVG(locations)
	return lake
}

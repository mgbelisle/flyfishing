package flyfishing

import (
	"log"
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
func (l Lake) ShowCastLogs(castLogs []CastLog) {
	log.Println(len(castLogs))
}

func NewLake() Lake {
	lake := Lake{Length: 500, Width: 300}
	sweetSpot := lake.RandLoc()
	// Adds 1000 fish to the lake placed randomly around a sweet
	// spot.
	for i := 0; i < 1000 ; i++ {
		loc := lake.RandLoc()
		loc.X = sweetSpot.X + (loc.X - sweetSpot.X) / 4
		loc.Y = sweetSpot.Y + (loc.Y - sweetSpot.Y) / 4
		lake.fishes = append(lake.fishes, lake.newFish(loc))
	}
	return lake
}

package flyfishing

import (
	"fmt"
	"math/rand"
	"math"
)

// This interface for a fish makes the lake struct super flexible.
type Fish interface {
	getLocation() Location
	lureWith(Fly, Location) bool
}

// Trout is kind of like a class.  It can be used on anything that
// requires a Fish interface because it implements getLocation() and
// lureWith().
type Trout struct {
	location Location
}
func (t Trout) getLocation() Location {
	return t.location
}
func (t Trout) lureWith(fly Fly, loc Location) bool {
	return t.noticesFly(loc) && t.isHungry() && t.likesFlyType(fly)
}
func (t Trout) noticesFly(loc Location) bool {
	distance := math.Sqrt(math.Pow(loc.X - t.location.X, 2) +
		math.Pow(loc.Y - t.location.Y, 2))
	return rand.Float64() < 1 / (1 + math.Pow(distance, 2))
}
func (_ Trout) isHungry() bool {
	return rand.Float64() > 0.5
}
func (_ Trout) likesFlyType(fly Fly) bool {
	switch fly.(type) {
	case Caddis: return true
	case WoollyBugger: return true
	}
	return false
}
func (t Trout) String() string {
	return fmt.Sprintf("%T at (%f, %f)", t, t.location.X, t.location.Y)
}

// Cutthroat is kind of like a subclass of Trout.  It likes different
// flies than a normal trout.
type Cutthroat struct { Trout }
func (_ Cutthroat) likesFlyType(fly Fly) bool {
	switch fly.(type) {
	case Caddis: return true
	case ParachuteAdams: return true
	}
	return false
}

// Rainbow is kind of like a subclass of Trout.  Notice how it is less
// hungry than a normal Trout.
type Rainbow struct { Trout }
func (_ Rainbow) isHungry() bool {
	return rand.Float64() > 0.3
}

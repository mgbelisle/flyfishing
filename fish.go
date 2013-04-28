package flyfishing

import (
	"fmt"
	"math/rand"
	"math"
)

// This interface for a fish makes the lake struct super flexible.
type Fish interface {
	Loc() Location
	LureWith(fly Fly, distance float64) bool
}

// Trout is kind of like a class
type Trout struct {
	loc Location
}
func (t Trout) Loc() Location {
	return t.loc
}
func (t Trout) LureWith(fly Fly, distance float64) bool {
	return t.noticesFly(distance) && t.isHungry() && t.likesFlyType(fly)
}
func (_ Trout) noticesFly(distance float64) bool {
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
	return fmt.Sprintf("%T at (%f, %f)", t, t.loc.X, t.loc.Y)
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

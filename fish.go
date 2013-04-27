package flyfishing

import (
	"math/rand"
	"math"
)

// This interface for a fish makes the lake struct super flexible.
type Fish interface {
	Loc() Location
	LureWith(fly Fly, distance float64) bool
}

// trout is kind of like a class
type trout struct {
	loc Location
}
func (t trout) Loc() Location {
	return t.loc
}
func (t trout) LureWith(fly Fly, distance float64) bool {
	return t.noticesFly(distance) && t.isHungry() && t.likesFlyType(fly)
}
func (_ trout) noticesFly(distance float64) bool {
	return 1 / (1 + math.Pow(distance, 2)) > rand.Float64()
}
func (_ trout) isHungry() bool {
	return rand.Float64() > 0.5
}
func (_ trout) likesFlyType(fly Fly) bool {
	switch fly.(type) {
	case Caddis: return true
	case ParachuteAdams: return true
	case WoollyBugger: return true
	}
	return false
}

// Cutthroat is kind of like a subclass of trout.  Notice how it is
// pickier about what types of flies it likes.
type Cutthroat struct { trout }
func (_ Cutthroat) likesFlyType(fly Fly) bool {
	switch fly.(type) {
	case Caddis: return true
	case ParachuteAdams: return true
	}
	return false
}

// Rainbow is kind of like a subclass of trout.  Notice how it is less
// hungry than a normal trout.
type Rainbow struct { trout }
func (_ Rainbow) isHungry() bool {
	return rand.Float64() > 0.3
}

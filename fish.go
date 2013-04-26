package flyfishing

import "math/rand"

type Fish interface {
	XLoc, YLoc float
	likesFlyType(Fly) bool
}

func (f Fish) LureWith(fly Fly, distance float) bool {
	return f.noticesFly(distance) && f.isHungry() && f.likesFlyType(fly)
}

func (f Fish) isHungry() bool {
	return rand.Float32() > 0.5
}

func (f Fish) noticesFly(distance float) bool {
	return 1 / (1 + distance ** 2) > rand.Float32()
}

type Trout struct {
	XLoc, YLoc float
}

type Cutthroat struct {
	Trout
}

func (f Cutthroat) likesFlyType(fly Fly) bool {
	switch fly.(type) {
	case Caddis: return true
	case ParachuteAdams: return true
	}
	return false
}

type Rainbow struct {
	Trout
}

func (f Rainbow) likesFlyType(fly Fly) bool {
	switch fly.(type) {
	case ParachuteAdams: return true
	case WoollyBugger: return true
	}
	return false
}

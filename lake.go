package flyfishing

type Lake struct {
	XBound float
	YBound float
	Fishes []Fish
}

func (l Lake) CastFly(fly Fly, x float, y float) Fish {
	for _, fish := range l.Fishes {
		
	}
}
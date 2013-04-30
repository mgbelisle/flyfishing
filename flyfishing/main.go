package main

import (
	"flyfishing"
)

func main() {
	lake := flyfishing.NewLake()
	biteLocations := castNTimesAsync(1000, lake)
	lake.LocationsToSVG(biteLocations)
}

func castNTimes(n int, lake flyfishing.Lake) []flyfishing.Location {
	locations := []flyfishing.Location{}
	for i := 0; i < n; i++ {
		loc := lake.RandLoc()
		fly := flyfishing.Caddis{}
		fish := lake.CastInto(fly, loc)
		if fish != nil {
			locations = append(locations, loc)
		}
	}
	return locations
}

func castNTimesAsync(n int, lake flyfishing.Lake) []flyfishing.Location {
	locations := []flyfishing.Location{}
	done := make(chan bool)
	for i := 0; i < n; i++ {
		go func() {
			loc := lake.RandLoc()
			fly := flyfishing.Caddis{}
			fish := lake.CastInto(fly, loc)
			if fish != nil {
				locations = append(locations, loc)
			}
			done <- true
		}()
	}
	for i := 0; i < n; i++ {
		<-done
	}
	return locations
}

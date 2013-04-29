package main

import (
	"flyfishing"
)

func main() {
	lake := flyfishing.NewLake()
	biteLocations := castNTimes(10, lake)
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
	locationChan := make(chan flyfishing.Location)
	for i := 0; i < n; i++ {
		go func() {
			loc := lake.RandLoc()
			fly := flyfishing.Caddis{}
			fish := lake.CastInto(fly, loc)
			if fish != nil {
				locationChan <- loc
			}
		}()
	}
	locations := []flyfishing.Location{}
	for {
		locations = append(locations, <-locationChan)
	}
	return locations
}

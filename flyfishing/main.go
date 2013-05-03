package main // Naming a package main means that nothing else can
	     // import it and it is compiled into an executable.

import (
	"flyfishing" // Custom package for this demo
	"fmt"
)

// main() executes first.  It instantiates a lake with a bunch of
// fish, does 5000 casts into the lake, and prints a map of successful
// cast locations.
func main() {
	lake := flyfishing.NewLake()
	biteLocations := castNTimesAsync(5000, lake)
	svg := lake.LocationsToSVG(biteLocations)
	fmt.Println("Map created:", svg.Name())
}

// Casts into the lake n times, returning the locations where fish
// were caught
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

// Casts into the lake n times in parallel, returning the locations
// where the fish were caught
func castNTimesAsync(n int, lake flyfishing.Lake) []flyfishing.Location {
	castLogChan := make(chan flyfishing.CastLog)
	for i := 0; i < n; i++ {
		// The go keyword means execute this function in
		// another goroutine.
		go func() {
			loc := lake.RandLoc()
			fly := flyfishing.Caddis{}
			fish := lake.CastInto(fly, loc)
			castLogChan <- flyfishing.CastLog{loc, fish}
		}()
	}
	locations := []flyfishing.Location{}
	for i := 0; i < n; i++ {
		castLog := <-castLogChan
		if castLog.Fish != nil {
			locations = append(locations, castLog.Location)
		}
	}
	return locations
}

// Naming a package main means that nothing else can import it and it
// is compiled into an executable.
package main

import (
	"github.com/mgbelisle/flyfishing" // Custom package for this demo
	"io"
	"os"
)

// main() executes first.  It instantiates a lake with a bunch of
// fish, does n casts into the lake, and prints a map of where fish
// were caught.
func main() {
	lake := flyfishing.NewLake()
	locations := castNTimes(3000, lake)
	io.Copy(os.Stdout, lake.LocationsToSVG(locations))
}

// Casts into the lake n times, returning the locations where fish
// were caught
func castNTimes(n int, lake flyfishing.Lake) []flyfishing.Location {
	locations := []flyfishing.Location{}
	for i := 0; i < n; i++ {
		location := lake.RandLocation()
		fly := flyfishing.Caddis{}
		fish := lake.CastInto(fly, location)
		if fish != nil {
			locations = append(locations, location)
		}
	}
	return locations
}

// Casts into the lake n times in parallel, returning the locations
// where the fish were caught.
func castNTimesAsync(n int, lake flyfishing.Lake) []flyfishing.Location {
	// Kicks off one goroutine for each cast.
	catchChan := make(chan bool)
	locationChan := make(chan flyfishing.Location)
	for i := 0; i < n; i++ {
		// This is an inline function, and go executes it in a
		// new goroutine via the go keyword.
		go func() {
			location := lake.RandLocation()
			fly := flyfishing.Caddis{}
			fish := lake.CastInto(fly, location)
			if fish == nil {
				catchChan <- false
			} else {
				catchChan <- true
				locationChan <- location
			}
		}()
	}

	// Waits for all the goroutines to finish, collecting the
	// locations for the successful ones.
	locations := []flyfishing.Location{}
	for i := 0; i < n; i++ {
		if <-catchChan {
			locations = append(locations, <-locationChan)
		}
	}
	return locations
}

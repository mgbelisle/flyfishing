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
	locations := castNTimesAsync(5000, lake)
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
// where the fish were caught
func castNTimesAsync(n int, lake flyfishing.Lake) []flyfishing.Location {
	locations := []flyfishing.Location{}
	c := make(chan bool)
	for i := 0; i < n; i++ {
		// This is an inline function, and go executes it in a
		// new goroutine via the go keyword.
		go func() {
			location := lake.RandLocation()
			fly := flyfishing.Caddis{}
			fish := lake.CastInto(fly, location)
			if fish != nil {
				locations = append(locations, location)
			}
			c <- true
		}()
	}
	for i := 0; i < n; i++ {
		<-c
	}
	return locations
}

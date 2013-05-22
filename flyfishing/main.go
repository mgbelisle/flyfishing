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
	castLogs := castNTimesAsync(5000, lake)
	io.Copy(os.Stdout, lake.CastLogsToSVG(castLogs))
}

// Casts into the lake n times, returning the locations where fish
// were caught
func castNTimes(n int, lake flyfishing.Lake) []flyfishing.CastLog {
	castLogs := []flyfishing.CastLog{}
	for i := 0; i < n; i++ {
		loc := lake.RandLoc()
		fly := flyfishing.Caddis{}
		fish := lake.CastInto(fly, loc)
		castLog := flyfishing.CastLog{Location: loc, Fly: fly, Fish: fish}
		castLogs = append(castLogs, castLog)
	}
	return castLogs
}

// Casts into the lake n times in parallel, returning the locations
// where the fish were caught
func castNTimesAsync(n int, lake flyfishing.Lake) []flyfishing.CastLog {
	// A CastLog channel is like a pipe for CastLog objects.
	castLogChan := make(chan flyfishing.CastLog)
	for i := 0; i < n; i++ {
		// The go keyword means execute this function in
		// another goroutine.
		go castOnce(lake, castLogChan)
	}
	castLogs := []flyfishing.CastLog{}
	for i := 0; i < n; i++ {
		// <-castLogChan pulls a CastLog object out of the
		// channel, waiting if necessary.
		castLogs = append(castLogs, <-castLogChan)
	}
	return castLogs
}

func castOnce(lake flyfishing.Lake, castLogChan chan flyfishing.CastLog) {
	loc := lake.RandLoc()
	fly := flyfishing.Caddis{}
	fish := lake.CastInto(fly, loc)
	castLog := flyfishing.CastLog{Location: loc, Fly: fly, Fish: fish}
	// This sends a CastLog object into the channel.
	castLogChan <- castLog
}

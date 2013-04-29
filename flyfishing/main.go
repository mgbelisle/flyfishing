package main

import (
	"flyfishing"
)

func main() {
	lake := flyfishing.NewLake()
	castLogs := castNTimesAsync(100, lake)
	lake.ShowCastLogs(castLogs)
}

func castNTimes(n int, lake flyfishing.Lake) []flyfishing.CastLog {
	castLogs := []flyfishing.CastLog{}
	for i := 0; i < n; i++ {
		loc := lake.RandLoc()
		fish := lake.CastInto(flyfishing.Caddis{}, lake.RandLoc())
		castLogs = append(castLogs, flyfishing.CastLog{loc, fish})
	}
	return castLogs
}

func castNTimesAsync(n int, lake flyfishing.Lake) []flyfishing.CastLog {
	castChan := make(chan flyfishing.CastLog)
	for i := 0; i < n; i++ {
		go func() {
			loc := lake.RandLoc()
			fish := lake.CastInto(flyfishing.Caddis{}, lake.RandLoc())
			castChan <- flyfishing.CastLog{loc, fish}
		}()
	}
	castLogs := []flyfishing.CastLog{}
	for i := 0; i < n; i++ {
		castLogs = append(castLogs, <-castChan)
	}
	return castLogs
}

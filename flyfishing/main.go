package main

import (
	"flyfishing"
)

func main() {
	lake := flyfishing.NewLake()
	castLogs := castNTimesAsync(1000, lake)
	lake.ShowCastLogs(castLogs)
}

func castNTimes(n int, lake flyfishing.Lake) []flyfishing.CastLog {
	castLogs := []flyfishing.CastLog{}
	for i := 0; i < n; i++ {
		loc := lake.RandLoc()
		fly := flyfishing.Caddis{}
		fish := lake.CastInto(fly, lake.RandLoc())
		castLog := flyfishing.CastLog{loc, fly, fish}
		castLogs = append(castLogs, castLog)
	}
	return castLogs
}

func castNTimesAsync(n int, lake flyfishing.Lake) []flyfishing.CastLog {
	castLogChan := make(chan flyfishing.CastLog)
	for i := 0; i < n; i++ {
		go func() {
			loc := lake.RandLoc()
			fly := flyfishing.Caddis{}
			fish := lake.CastInto(fly, lake.RandLoc())
			castLog := flyfishing.CastLog{loc, fly, fish}
			castLogChan <- castLog
		}()
	}
	castLogs := []flyfishing.CastLog{}
	for i := 0; i < n; i++ {
		castLogs = append(castLogs, <-castLogChan)
	}
	return castLogs
}

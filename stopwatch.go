package main

import "time"

type Stopwatch struct {
	timeDate time.Time
	isGoing  bool
	timeMark []time.Time
}

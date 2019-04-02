package main

import "math/rand"

//Sector - содержит о себе всю информацию о вселенной.
type Sector struct {
	hexCoordinates string
	zoneType       string
}

func roll1dX(x int, mod int) int {
	if x < 1 {
		x = 1
	}
	return randInt(1, x) + mod
}

func randInt(min int, max int) int {
	return min + rand.Intn(max)
}

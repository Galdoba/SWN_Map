package main

import "github.com/Galdoba/utils"

type Star struct {
	Name        string
	starColor   string
	class       string
	temperature float64
	mass        float64
	description string
}

func NewStar(clss string) *Star {
	star := Star{}
	star.class = clss
	switch clss {
	case "O":
		star.starColor = "Bright Blue"
		star.temperature = utils.RandFloat(33.0, 99.0, 3)
		star.mass = utils.RandFloat(16.0, 50.0, 2)
	case "B":
		star.starColor = "Blue-White"
		star.temperature = utils.RandFloat(10.0, 33.0, 3)
		star.mass = utils.RandFloat(2.1, 16.0, 2)
	case "A":
		star.starColor = "White"
		star.temperature = utils.RandFloat(7.5, 10.0, 3)
		star.mass = utils.RandFloat(1.4, 2.1, 2)
	case "F":
		star.starColor = "Yellowish-White"
		star.temperature = utils.RandFloat(6, 7.5, 3)
		star.mass = utils.RandFloat(1.04, 1.4, 2)
	case "G":
		star.starColor = "Yellow"
		star.temperature = utils.RandFloat(5.2, 6, 3)
		star.mass = utils.RandFloat(0.8, 1.04, 2)
	case "K":
		star.starColor = "Orange"
		star.temperature = utils.RandFloat(3.7, 5.2, 3)
		star.mass = utils.RandFloat(0.45, 0.8, 2)
	case "M":
		star.starColor = "Red"
		star.temperature = utils.RandFloat(3, 3.7, 3)
		star.mass = utils.RandFloat(0.35, 0.45, 2)
	case "C":
		star.starColor = "Dim Red"
		star.temperature = utils.RandFloat(2.7, 3, 3)
		star.mass = utils.RandFloat(0.26, 0.35, 2)
		star.description = "Red Giant (Swollen Dying Star)"
	case "T":
		star.starColor = "Faint Brown"
		star.temperature = utils.RandFloat(2.3, 2.7, 3)
		star.mass = utils.RandFloat(0.17, 0.26, 2)
		star.description = "Brown Dwarf (Tiny Proto Star)"
	case "D":
		star.starColor = "Faint White"
		star.temperature = utils.RandFloat(1.6, 2.3, 3)
		star.mass = utils.RandFloat(0.14, 0.17, 2)
		star.description = "White Dwarf (Star Remnant)"
	case "Black Hole":
		star.starColor = "Black"
		star.temperature = utils.RandFloat(0, 0, 3)
		star.mass = utils.RandFloat(0, 0, 2)
		star.description = "Black Hole with companion orange star it is being drawn into its gravity well. Powerful tidal forces make this system extremly Hazardous. The system contains the broken remnants of planets."
	case "Neutron Star":
		star.starColor = "White"
		star.temperature = utils.RandFloat(0, 0, 3)
		star.mass = utils.RandFloat(0, 0, 2)
		star.description = "Neutron Star, a tiny collapsed remnant of a star that went supernova. It exhibits strong gravity and radioactivity). Few remnants of the star’s system survived its nova."

	default:
		star.starColor = ""
		star.temperature = utils.RandFloat(0, 0, 3)
		star.mass = utils.RandFloat(0, 0, 2)
		star.description = "UNKNOWN TYPE"
	}
	return &star
}

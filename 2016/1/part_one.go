package one

import (
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type CardinalDirection struct {
	north bool
	east  bool
	south bool
	west  bool
}

func SetCardinalDirection(letter string, cardinalDirection CardinalDirection) CardinalDirection {
	if cardinalDirection.east == true {
		if letter == "R" {
			cardinalDirection.south = true
		} else {
			cardinalDirection.north = true
		}
		cardinalDirection.east = false
	} else if cardinalDirection.west == true {
		if letter == "R" {
			cardinalDirection.north = true
		} else {
			cardinalDirection.south = true
		}
		cardinalDirection.west = false
	} else if cardinalDirection.north == true {
		if letter == "R" {
			cardinalDirection.east = true
		} else {
			cardinalDirection.west = true
		}
		cardinalDirection.north = false
	} else if cardinalDirection.south == true {
		if letter == "R" {
			cardinalDirection.west = true
		} else {
			cardinalDirection.east = true
		}
		cardinalDirection.south = false
	}

	return cardinalDirection
}

func GetNewXAndY(cardinalDirection CardinalDirection, coordinate Coordinate, x, y, steps int) (int, int) {
	if cardinalDirection.north == true {
		y = coordinate.y + steps
	} else if cardinalDirection.east == true {
		x = coordinate.x + steps
	} else if cardinalDirection.south == true {
		y = coordinate.y - steps
	} else if cardinalDirection.west == true {
		x = coordinate.x - steps
	}

	return x, y
}

func StepsFromStart(coordinate Coordinate) int {
	return int(math.Abs(float64(coordinate.x)) + math.Abs(float64(coordinate.y)))
}

func PartOne(directions string) int {
	cardinalDirection := CardinalDirection{
		north: false,
		east:  false,
		south: false,
		west:  false,
	}

	x := 0
	y := 0

	coordinate := Coordinate{x, y}

	parts := strings.Split(directions, ", ")
	for i, direction := range parts {
		letterAndSteps := strings.Split(direction, "")
		letter := letterAndSteps[0]
		steps, err := strconv.Atoi(strings.Join(letterAndSteps[1:], ""))

		if err == nil {
			if i == 0 {
				if letter == "R" {
					cardinalDirection.east = true
				} else {
					cardinalDirection.west = true
				}
			} else {
				cardinalDirection = SetCardinalDirection(letter, cardinalDirection)
			}
		}

		x, y = GetNewXAndY(cardinalDirection, coordinate, x, y, steps)

		coordinate.x = x
		coordinate.y = y
	}

	return StepsFromStart(coordinate)
}

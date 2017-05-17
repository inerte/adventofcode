package one

import (
	"strconv"
	"strings"
)

func coordinateHasBeenWalked(visitedCoordinates map[Coordinate]struct{}, coordinate Coordinate) bool {
	_, ok := visitedCoordinates[coordinate]
	return ok
}

func PartTwo(directions string) int {
	cardinalDirection := CardinalDirection{
		north: false,
		east:  false,
		south: false,
		west:  false,
	}

	x := 0
	y := 0

	coordinate := Coordinate{x, y}

	visitedCoordinates := make(map[Coordinate]struct{})

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

		if x > coordinate.x {
			// http://stackoverflow.com/questions/10485743/contains-method-for-a-slice#comment13552187_10486196
			for i = coordinate.x + 1; i <= x; i++ {
				newCoordinate := Coordinate{i, coordinate.y}
				if coordinateHasBeenWalked(visitedCoordinates, newCoordinate) {
					return StepsFromStart(newCoordinate)
				}
				visitedCoordinates[newCoordinate] = struct{}{}
			}
		}
		if x < coordinate.x {
			for i = coordinate.x - 1; i >= x; i-- {
				newCoordinate := Coordinate{i, coordinate.y}
				if coordinateHasBeenWalked(visitedCoordinates, newCoordinate) {
					return StepsFromStart(newCoordinate)
				}
				visitedCoordinates[newCoordinate] = struct{}{}
			}
		}

		if y > coordinate.y {
			for i = coordinate.y + 1; i <= y; i++ {
				newCoordinate := Coordinate{coordinate.x, i}
				if coordinateHasBeenWalked(visitedCoordinates, newCoordinate) {
					return StepsFromStart(newCoordinate)
				}
				visitedCoordinates[newCoordinate] = struct{}{}
			}
		}
		if y < coordinate.y {
			for i = coordinate.y - 1; i >= y; i-- {
				newCoordinate := Coordinate{coordinate.x, i}
				if coordinateHasBeenWalked(visitedCoordinates, newCoordinate) {
					return StepsFromStart(newCoordinate)
				}
				visitedCoordinates[newCoordinate] = struct{}{}
			}
		}

		coordinate.x = x
		coordinate.y = y
	}

	return StepsFromStart(coordinate)
}

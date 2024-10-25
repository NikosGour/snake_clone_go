package main

import (
	"os"

	log "github.com/NikosGour/logging/src"
)

type Direction uint8

const (
	Direction_UP Direction = iota
	Direction_RIGHT
	Direction_DOWN
	Direction_LEFT
	_Direction_Count
)

func (dir Direction) String() string {
	if dir >= _Direction_Count {
		log.Error("Passed Direction value of %d when the max %d", dir, _Direction_Count-1)
		os.Exit(1)
	}
	switch dir {
	case Direction_UP:
		return "UP"
	case Direction_RIGHT:
		return "RIGHT"
	case Direction_DOWN:
		return "DOWN"
	case Direction_LEFT:
		return "LEFT"
	}
	// Only possible if I add new Direction in the enum and no in the switch
	log.Error("You forgot to implement String() for one or more Directions")
	os.Exit(1)
	return ""
}

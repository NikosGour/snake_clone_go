package main

import (
	"snake_clone/src/build"
)

func main() {
	game := newGame(build.DEBUG_MODE)
	game.runGameLoop()
}

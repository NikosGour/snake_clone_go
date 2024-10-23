package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Fruit struct {
	game_ctx *Game

	x     int
	y     int
	color rl.Color
}

func newFruit(game_ctx *Game) *Fruit {
	this := Fruit{x: 5, y: 5, color: rl.Orange}
	this.game_ctx = game_ctx
	return &this
}

func (this *Fruit) draw() {
	this.game_ctx.grid.drawCell(this.x, this.y, this.color)
}

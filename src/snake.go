package main

import (
	"fmt"
	"strings"

	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Snake struct {
	game_ctx *Game

	body  []rl.Vector2
	head  *rl.Vector2
	tail  *rl.Vector2
	color rl.Color

	direction Direction
}

type Direction uint8

const (
	Direction_UP Direction = iota
	Direction_RIGHT
	Direction_DOWN
	Direction_LEFT
)

func newSnake(game_ctx *Game) *Snake {
	this := Snake{color: rl.Green, direction: Direction_RIGHT}
	this.game_ctx = game_ctx

	this.body = make([]rl.Vector2, 0, Grid_columns*Grid_rows)
	this.addBodyPart(0, 0)
	this.addBodyPart(0, 1)
	this.addBodyPart(1, 1)

	this.setHead()
	this.setTail()
	return &this
}

func (this *Snake) draw() {
	for i := range this.body {
		body_part := &this.body[i]
		color := rl.Black
		// log.Debug("body_part: %p, tail: %p, head: %p", body_part, this.tail, this.head)
		if body_part == this.tail {
			color.R += 0xFF
		}
		if body_part == this.head {
			color.B += 0xFF
		}

		if color == rl.Black {
			color = this.color
		}

		this.game_ctx.grid.drawCell(int(body_part.X), int(body_part.Y), color)

	}
}

func (this *Snake) addBodyPart(x int, y int) {
	this.body = append(this.body, rl.NewVector2(float32(x), float32(y)))
}

func (this *Snake) setTail() {
	this.tail = &this.body[0]
}

func (this *Snake) setHead() {
	this.head = &this.body[len(this.body)-1]
}

func (this *Snake) print() {
	for i := range this.body {
		body_part := &this.body[i]
		log.Debug("Part %d: %p", i, body_part)
	}
	log.Debug("Head: %p", this.head)
	log.Debug("Tail: %p", this.tail)
	var str strings.Builder

	str.WriteString("snake: {\n")
	indent := strings.Repeat("\t", 6)
	for _, body_part := range this.body {
		coords := fmt.Sprintf(indent+"x: %d, y: %d\n", int(body_part.X), int(body_part.Y))
		str.WriteString(coords)
	}
	head := fmt.Sprintf(indent+"head: {x: %d, y: %d}\n", int(this.head.X), int(this.head.Y))
	str.WriteString(head)

	tail := fmt.Sprintf(indent+"tail: {x: %d, y: %d}\n", int(this.tail.X), int(this.tail.Y))
	str.WriteString(tail)
	str.WriteString(indent + "}")

	log.Debug("%s", str.String())

}

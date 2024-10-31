package main

import (
	"errors"
	"fmt"
	"slices"
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

func newSnake(game_ctx *Game) *Snake {
	this := Snake{color: rl.Green, direction: Direction_RIGHT}
	this.game_ctx = game_ctx
	fmt.Println(this.direction)

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
		if this.game_ctx.debug_mode {
			if body_part == this.tail {
				color.R += 0xFF
			}
			if body_part == this.head {
				color.B += 0xFF
			}
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

var (
	ErrorOutOfBoundsUp    = errors.New("OutOfBoundsUp")
	ErrorOutOfBoundsDown  = errors.New("OutOfBoundsDown")
	ErrorOutOfBoundsLeft  = errors.New("OutOfBoundsLeft")
	ErrorOutOfBoundsRight = errors.New("OutOfBoundsRight")
	ErrorHitBody          = errors.New("HitBody")
)

func (this *Snake) move() error {
	this.setHead()
	switch this.direction {
	case Direction_UP:
		if this.head.Y == 0 {
			return fmt.Errorf("%s at {x: %d, y: %d}", ErrorOutOfBoundsUp, int(this.head.X), int(this.head.Y))
		}
	case Direction_DOWN:
		if this.head.Y == Grid_rows-1 {
			return fmt.Errorf("%s at {x: %d, y: %d}", ErrorOutOfBoundsDown, int(this.head.X), int(this.head.Y))
		}
	case Direction_LEFT:
		if this.head.X == 0 {
			return fmt.Errorf("%s at {x: %d, y: %d}", ErrorOutOfBoundsLeft, int(this.head.X), int(this.head.Y))
		}
	case Direction_RIGHT:
		if this.head.X == Grid_columns-1 {
			return fmt.Errorf("%s at {x: %d, y: %d}", ErrorOutOfBoundsRight, int(this.head.X), int(this.head.Y))
		}
	}
	var move_v rl.Vector2

	switch this.direction {
	case Direction_UP:
		move_v = rl.NewVector2(this.head.X, this.head.Y-1)
	case Direction_DOWN:
		move_v = rl.NewVector2(this.head.X, this.head.Y+1)
	case Direction_LEFT:
		move_v = rl.NewVector2(this.head.X-1, this.head.Y)
	case Direction_RIGHT:
		move_v = rl.NewVector2(this.head.X+1, this.head.Y)
	}
	if slices.Contains(this.body, move_v) {
		return fmt.Errorf("%s going %d head at {x: %d, y: %d}, body segment at {x: %d, y: %d}",
			ErrorHitBody,
			this.direction,
			int(this.head.X),
			int(this.head.Y),
			int(move_v.X),
			int(move_v.Y),
		)
	}

	this.addBodyPart(int(move_v.X), int(move_v.Y))
	this.body = slices.Delete(this.body, 0, 1)
	this.setHead()
	this.setTail()
	return nil
}

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Grid_columns = 60
	Grid_rows    = 32
)

type Grid struct {
	game_ctx *Game

	cell_width   int
	cell_padding int
	cell_size    rl.Vector2
	cell_color   rl.Color

	starting_point rl.Vector2
}

func newGrid(game_ctx *Game) *Grid {
	this := Grid{cell_width: 20, cell_padding: 5, cell_color: rl.NewColor(0x33, 0x33, 0x33, 0xFF)}

	this.game_ctx = game_ctx
	this.cell_size = rl.NewVector2(float32(this.cell_width), float32(this.cell_width))

	this.starting_point = rl.NewVector2(
		float32((this.game_ctx.screen_width-(this.cell_width+this.cell_padding)*Grid_columns)/2),
		float32((this.game_ctx.screen_height-(this.cell_width+this.cell_padding)*Grid_rows)/2),
	)

	return &this

}

func (this *Grid) draw() {

	copy_starting_point := this.starting_point
	for range Grid_rows {
		point := copy_starting_point

		for range Grid_columns {
			rl.DrawRectangleV(
				point,
				this.cell_size,
				this.cell_color,
			)

			point.X += float32(this.cell_width + this.cell_padding)
			copy_starting_point.X += float32(this.cell_width + this.cell_padding)
		}
		copy_starting_point.X = this.starting_point.X
		copy_starting_point.Y += float32(this.cell_width + this.cell_padding)
	}

}

func (this *Grid) drawCell(x int, y int, color rl.Color) {
	draw_vec := rl.NewVector2(
		float32(int(this.starting_point.X)+(this.cell_width+this.cell_padding)*x),
		float32(int(this.starting_point.Y)+(this.cell_width+this.cell_padding)*y),
	)

	rl.DrawRectangleV(draw_vec, this.cell_size, color)

}

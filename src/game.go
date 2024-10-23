package main

import (
	log "github.com/NikosGour/logging/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	current_monitor int
	monitor_height  int
	monitor_width   int
	screen_height   int
	screen_width    int

	grid  *Grid
	snake *Snake
	fruit *Fruit
}

func newGame() Game {

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 600, "Snake Clone by Nikos Gournakis")

	rl.SetTargetFPS(60)

	if !rl.IsWindowMaximized() {
		rl.MaximizeWindow()
	}

	this := Game{}
	log.Debug("%#v", this)

	return this
}

func (this *Game) init() {
	this.grid = newGrid(this)
	this.snake = newSnake(this)
	this.fruit = newFruit(this)

}

func (this *Game) runGameLoop() {
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		this.configureMonitorScreenSizes()
		// log.Debug("monitor: %v, monitor_height: %v, monitor_width: %v", this.current_monitor, this.monitor_height, this.monitor_width)
		// log.Debug("screen_height: %v, screen_width: %v", this.screen_height, this.screen_width)

		if this.screen_height > 600 && this.screen_width > 800 {
			this.init()
		} else if this.monitor_height == 600 || this.screen_width == 800 {
			this.init()
		}

		rl.BeginDrawing()
		// ---------------- DRAWING ----------------------------
		if this.grid != nil {
			this.grid.draw()
			this.snake.draw()
			this.snake.print()
			this.fruit.draw()
		}

		rl.ClearBackground(rl.NewColor(0x18, 0x18, 0x18, 0xFF))
		// ---------------- END DRAWING ------------------------
		rl.EndDrawing()
	}
}

func (this *Game) configureMonitorScreenSizes() {
	this.current_monitor = rl.GetCurrentMonitor()
	this.monitor_height = rl.GetMonitorHeight(this.current_monitor)
	this.monitor_width = rl.GetMonitorWidth(this.current_monitor)

	this.screen_height = rl.GetScreenHeight()
	this.screen_width = rl.GetScreenWidth()
}

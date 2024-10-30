package main

import (
	"os"
	"time"

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

	event_ticker *time.Ticker

	debug_mode bool
}

func newGame(debug bool) Game {

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 600, "Snake Clone by Nikos Gournakis")

	rl.SetTargetFPS(60)

	if !rl.IsWindowMaximized() {
		rl.MaximizeWindow()
	}

	this := Game{debug_mode: debug}
	log.Debug("%#v", this)

	return this
}

func (this *Game) init() {
	this.grid = newGrid(this)
	this.snake = newSnake(this)
	this.fruit = newFruit(this)
	this.event_ticker = time.NewTicker(time.Second / 3)

}

func (this *Game) runGameLoop() {
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		this.configureMonitorScreenSizes()
		// log.Debug("monitor: %v, monitor_height: %v, monitor_width: %v", this.current_monitor, this.monitor_height, this.monitor_width)
		// log.Debug("screen_height: %v, screen_width: %v", this.screen_height, this.screen_width)

		// Triggering events
		if this.grid != nil {
			if rl.IsKeyPressed(rl.KeyW) && this.snake.direction != Direction_DOWN {
				this.snake.direction = Direction_UP
			} else if rl.IsKeyPressed(rl.KeyS) && this.snake.direction != Direction_UP {
				this.snake.direction = Direction_DOWN
			} else if rl.IsKeyPressed(rl.KeyD) && this.snake.direction != Direction_LEFT {
				this.snake.direction = Direction_RIGHT
			} else if rl.IsKeyPressed(rl.KeyA) && this.snake.direction != Direction_RIGHT {
				this.snake.direction = Direction_LEFT
			}

			// Update state
			select {
			case <-this.event_ticker.C:
				this.update()
			default:
			}
		}

		if this.screen_height > 600 && this.screen_width > 800 && this.grid == nil {
			this.init()
		} else if (this.monitor_height == 600 || this.monitor_width == 800) && this.grid == nil {
			this.init()
		}

		rl.BeginDrawing()
		// ---------------- DRAWING ----------------------------
		if this.grid != nil {
			this.grid.draw()
			this.snake.draw()
			this.fruit.draw()
		}

		rl.ClearBackground(rl.NewColor(0x18, 0x18, 0x18, 0xFF))
		// ---------------- END DRAWING ------------------------
		rl.EndDrawing()
	}
}

func (this *Game) update() {
	this.snake.print()
	err := this.snake.move()
	if err != nil {
		log.Error("Went out of bounds: %s", err)
		os.Exit(1)
	}
}

func (this *Game) configureMonitorScreenSizes() {
	this.current_monitor = rl.GetCurrentMonitor()
	this.monitor_height = rl.GetMonitorHeight(this.current_monitor)
	this.monitor_width = rl.GetMonitorWidth(this.current_monitor)

	this.screen_height = rl.GetScreenHeight()
	this.screen_width = rl.GetScreenWidth()
}

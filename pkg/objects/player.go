package objects

import (
	"zeroground/colors"

	"github.com/veandco/go-sdl2/sdl"
)

var size = int32(16)
var snakeBaseSpeed = uint64(100)
var snakeAccSpeed = uint64(50)

type player struct {
	snake
	AccOn      bool
	lastUpdate uint64
}

// func (p *Player) Position() sdl.Point {
// 	return p.snake.Position()
// }

// func (p *Player) Rect() []sdl.Rect {
// 	return p.snake.Rect()
// }

func (p *player) Handle(event sdl.Event) {
	var currdir Direction

	p.AccOn = false

	if event.GetType() == sdl.KEYDOWN {
		ke := event.(*sdl.KeyboardEvent)
		switch ke.Keysym.Scancode {
		case sdl.SCANCODE_UP:
			currdir = Dir_Up
		case sdl.SCANCODE_DOWN:
			currdir = Dir_Down
		case sdl.SCANCODE_LEFT:
			currdir = Dir_Left
		case sdl.SCANCODE_RIGHT:
			currdir = Dir_Right
		}
	}

	if currdir == Dir_Invalid {
		return
	}

	if p.direction() == currdir {
		p.AccOn = true
		return
	}

	p.SetDirection(currdir)
}

// func (p *Player) Direction() Direction {
// 	return p.snake.Direction()
// }

// func (p *Player) Draw(renderer *sdl.Renderer) {
// 	p.snake.Draw(renderer)
// }

func (p *player) Eat(food Food) {
	switch food.State() {
	case FoodFresh:
		p.grow()
	case FoodRotten:
		p.Alive = false
	default:
	}
}

func (p *player) Update() {
	speed := snakeBaseSpeed
	if p.AccOn {
		speed = snakeAccSpeed
	}
	if sdl.GetTicks64()-p.lastUpdate > speed {
		p.snake.Update()
		p.lastUpdate = sdl.GetTicks64()
	}
}

// func (p *Player) IsAlive() bool {
// 	return p.snake.IsAlive()
// }

// func (p *Player) Move() {
// 	p.snake.Move()
// }

// func (p *Player) SetAlive(a bool) {
// 	p.snake.SetAlive(a)
// }

// func (p *Player) Reset() {
// 	fmt.Println("Snake Reset")
// 	p.Snake.Reset()
// }

// func (p *Player) SetDirection(dir Direction) {
// 	p.snake.SetDirection(dir)
// }

func (s *player) Draw(renderer *sdl.Renderer) {
	for i, part := range s.Hitbox() {
		if i == 0 {
			renderer.SetDrawColor(colors.RGBA(colors.White()))
		} else {
			renderer.SetDrawColor(s.color.R, s.color.G, s.color.B, sdl.ALPHA_TRANSPARENT)
		}
		renderer.FillRect(part.BoundingRect())
		part.Draw(renderer)
		// f := part.Rect
		// f.H = 0
		// f == part.Rect ?? what do u think
	}
}

func NewPlayer(start sdl.Point, color sdl.Color) Player {
	p := &player{
		snake:      simpleSnake(start, color, 3, size),
		AccOn:      false,
		lastUpdate: sdl.GetTicks64(),
	}
	return p
}

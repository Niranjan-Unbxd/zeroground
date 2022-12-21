package objects

import (
	"github.com/veandco/go-sdl2/sdl"
)

const size = int32(16)

var playerMaxSpeed = uint64(50)

type player struct {
	*snake
	texture *sdl.Texture
	AccOn   bool
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
		speed = playerMaxSpeed
	}
	if sdl.GetTicks64()-p.lastUpdate > speed {
		p.snake.move()
		p.lastUpdate = sdl.GetTicks64()
	}
}

func (s *player) Draw(r *sdl.Renderer) {
	for i, part := range s.Hitbox() {
		rect := part.BoundingRect()
		if i == 0 {
			id := int32(s.direction() - 1)
			src := &sdl.Rect{X: id * size, Y: 0, W: size, H: size}
			r.Copy(s.texture, src, rect)
			// r.SetDrawColor(0, 95, 115, 0)
		} else {
			r.SetDrawColor(s.color.R, s.color.G, s.color.B, sdl.ALPHA_TRANSPARENT)
			r.FillRect(rect)
			r.SetDrawColor(148, 210, 189, 0)
			r.DrawRect(rect)
		}
		// f := part.Rect
		// f.H = 0
		// f == part.Rect ?? what do u think
	}
}

func NewPlayer(start sdl.Point, color sdl.Color, tex *sdl.Texture) Player {
	// snap to grid
	start.X = start.X - (start.X % size)
	start.Y = start.Y - (start.Y % size)

	p := &player{
		texture: tex,
		snake:   simpleSnake(start, color, 3, size),
		AccOn:   false,
	}
	return p
}

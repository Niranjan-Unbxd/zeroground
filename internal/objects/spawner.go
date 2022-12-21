package objects

import (
	"math/rand"
	"zeroground/pkg/colors"
	"zeroground/pkg/physics"
	"zeroground/pkg/ui"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	spawnerSize      = int32(16)
	spawnerSnakeSize = int32(8)
)

// Spawner spawns small enemy snakes
type Spawner struct {
	Pos   sdl.Point
	Color sdl.Color

	snake     *snake
	createdAt uint64
}

func (s *Spawner) Position() sdl.Point {
	return s.Pos
}

func (s *Spawner) Hitbox() []physics.Plane2D {
	var hbxs []physics.Plane2D

	p := physics.NewPlane2D(sdl.Rect{
		X: s.Pos.X - spawnerSize/2,
		Y: s.Pos.Y - spawnerSize/2,
		W: spawnerSize,
		H: spawnerSize,
	})
	hbxs = append(hbxs, *p)
	// hbxs = append(hbxs, physics.Plane2D{sdl.Rect{
	// 	X: s.Pos.X - spawnerSnakeSize/2,
	// 	Y: s.Pos.Y - spawnerSnakeSize/2,
	// 	W: spawnerSnakeSize,
	// 	H: spawnerSnakeSize,
	// }})

	hbxs = append(hbxs, s.snake.Hitbox()...)
	return hbxs
}

// func (s *Spawner) Rect() []sdl.Rect {
// 	rects := []sdl.Rect{
// 		{
// 			X: s.Pos.X - spawnerSize/2,
// 			Y: s.Pos.Y - spawnerSize/2,
// 			W: spawnerSize,
// 			H: spawnerSize,
// 		},
// 		{
// 			X: s.Pos.X - spawnerSnakeSize/2,
// 			Y: s.Pos.Y - spawnerSnakeSize/2,
// 			W: spawnerSnakeSize,
// 			H: spawnerSnakeSize,
// 		},
// 	}
// 	rects = append(rects, s.snake.Rect()...)
// 	return rects
// }

func (s *Spawner) Update() {
	if !s.snake.IsAlive() {
		s.snake.Reset()
	}

	speed := snakeBaseSpeed
	if sdl.GetTicks64()-s.snake.lastUpdate > speed {
		// try to change direction
		choice := Direction(rand.Int31n(4) + 1)
		s.snake.SetDirection(choice)
		// move spawned snake
		s.snake.move()
		s.snake.lastUpdate = sdl.GetTicks64()
	}
}

func (s *Spawner) Draw(r *sdl.Renderer) {
	rects := s.Hitbox()
	dark := colors.Darker(s.Color)
	// outside
	r.SetDrawColor(colors.RGBA(dark))
	r.FillRect(rects[0].BoundingRect())
	// inside
	// r2 := &sdl.Rect{
	// 	X: s.Pos.X - spawnerSnakeSize/2,
	// 	Y: s.Pos.Y - spawnerSnakeSize/2,
	// 	W: spawnerSnakeSize,
	// 	H: spawnerSnakeSize,
	// }
	// r.SetDrawColor(202, 103, 2, 0)
	// r.FillRect(r2)

	s.snake.Draw(r)
}

func (s *Spawner) Reset() {
	x := rand.Int31() % (ui.WindowWidth - 2*spawnerSize)
	y := rand.Int31() % (ui.WindowHeight - 2*spawnerSize)
	s.Pos = sdl.Point{X: x, Y: y}
	s.snake = simpleSnake(s.Pos, s.Color, 5, 8)
	s.createdAt = sdl.GetTicks64()
}

func NewSpawner(color sdl.Color) *Spawner {
	s := &Spawner{
		Color: color,
	}
	s.Reset()
	return s
}

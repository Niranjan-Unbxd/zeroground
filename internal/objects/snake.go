package objects

import (
	"zeroground/pkg/physics"
	"zeroground/pkg/ui"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	snakeBaseSpeed = uint64(100)
)

type snakeSegment struct {
	pos sdl.Point
	dir Direction
}

type snake struct {
	body []*snakeSegment
	len  int

	start sdl.Point
	color sdl.Color

	Idle          bool
	Alive         bool
	Size          int32
	InitialLength int32

	lastUpdate uint64

	// currState *State
	// states    map[string]*State
}

func (s *snake) Position() sdl.Point {
	return s.body[0].pos
}

func (s *snake) Hitbox() []physics.Plane2D {
	hbxs := make([]physics.Plane2D, 0, len(s.body))
	rect := sdl.Rect{X: 0, Y: 0, W: s.Size, H: s.Size}
	for _, part := range s.body {
		copied := rect
		copied.X = part.pos.X - s.Size/2
		copied.Y = part.pos.Y - s.Size/2
		hbxs = append(hbxs, physics.Plane2D{Rect: copied})
	}
	return hbxs
}

func (s *snake) HasIntersection(rects ...*sdl.Rect) (int, bool) {
	head := &sdl.Rect{
		X: s.body[0].pos.X - s.Size/2,
		Y: s.body[0].pos.Y - s.Size/2,
		W: s.Size,
		H: s.Size,
	}
	for i, r := range rects {
		if head.HasIntersection(r) {
			return i, true
		}
	}
	return -1, false
}

func (s *snake) Draw(r *sdl.Renderer) {
	for _, part := range s.Hitbox() {
		rect := part.BoundingRect()
		r.SetDrawColor(s.color.R, s.color.G, s.color.B, sdl.ALPHA_TRANSPARENT)
		r.FillRect(rect)
		// r.SetDrawColor(0, 18, 25, 0)
		// r.DrawRect(rect)
		r.SetDrawColor(238, 155, 0, 0)
		r.DrawRect(rect)
		// f := part.Rect
		// f.H = 0
		// f == part.Rect ?? what do u think
	}
}

func (s *snake) move() {
	tailPos := s.body[len(s.body)-1].pos
	for i := len(s.body) - 1; i > 0; i-- {
		s.body[i].pos = s.body[i-1].pos
	}

	// if can grow
	if s.len > len(s.body) {
		s.body = append(s.body, &snakeSegment{pos: tailPos})
	}

	head := s.body[0]
	switch head.dir {
	case Dir_Left:
		head.pos.X -= s.Size
	case Dir_Right:
		head.pos.X += s.Size
	case Dir_Up:
		head.pos.Y -= s.Size
	case Dir_Down:
		head.pos.Y += s.Size
	}

	// check if head collided with screen boundry
	hx := head.pos.X - s.Size/2
	hy := head.pos.Y - s.Size/2
	if hx < 0 || hy < 0 || hx+s.Size > ui.WindowWidth || hy+s.Size > ui.WindowHeight {
		s.Alive = false
	}

	// check if head collides with body
	for i := len(s.body) - 1; i > 0; i -= 1 {
		if head.pos.X == s.body[i].pos.X && head.pos.Y == s.body[i].pos.Y {
			s.Alive = false
		}
	}
}

func (s *snake) Update() {
	speed := snakeBaseSpeed
	if sdl.GetTicks64()-s.lastUpdate > speed {
		s.move()
		s.lastUpdate = sdl.GetTicks64()
	}
}

func (s *snake) IsAlive() bool {
	return s.Alive
}

// func (s *snake) SetAlive(a bool) {
// 	s.Alive = a
// }

func (s *snake) direction() Direction {
	return s.body[0].dir
}

func abs(i int32) int32 {
	if i < 0 {
		return -i
	}
	return i
}

func (s *snake) grow() {
	s.len += 1
}

func (s *snake) changeDirection(dir Direction) {
	s.body[0].dir = dir
}

func (s *snake) SetDirection(dir Direction) {
	currdir := s.direction()
	if currdir == dir {
		return
	}

	switch currdir {
	case Dir_Right:
		fallthrough
	case Dir_Left:
		if dir == Dir_Up || dir == Dir_Down {
			s.changeDirection(dir)
		}
	case Dir_Up:
		fallthrough
	case Dir_Down:
		if dir == Dir_Left || dir == Dir_Right {
			s.changeDirection(dir)
		}
	}
}

func (s *snake) Reset() {
	body := make([]*snakeSegment, 0, 10)
	for i := int32(0); i < s.InitialLength; i++ {
		pos := sdl.Point{X: s.start.X - i*s.Size, Y: s.start.Y}
		// fmt.Println(pos)
		body = append(body, &snakeSegment{pos: pos, dir: Dir_Right})
	}
	s.body = body
	s.len = len(s.body)
	s.SetDirection(Dir_Right)
	s.Alive = true
}

// func (s *snake) SetState(state string) {
// 	if s.currState != nil {
// 		s.currState.Clear()
// 	}
// 	s.currState = s.states[state]
// 	s.currState.Setup()
// }

func simpleSnake(start sdl.Point, color sdl.Color, len, size int32) *snake {
	s := &snake{
		start:         start,
		color:         color,
		Size:          size,
		Alive:         true,
		InitialLength: len,
		lastUpdate:    sdl.GetTicks64(),
	}
	s.Reset()
	return s
}

// type stateState struct{}
// NewSnake creates a new snake.
// length is snake's body length, snakes are always created with length >= 3
func NewSnake(start sdl.Point, color sdl.Color, length, size int32) Snake {
	if length < 3 {
		length = 3
	}
	s := simpleSnake(start, color, length, size)
	return s
}

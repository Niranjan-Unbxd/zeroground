package objects

import (
	"math/rand"
	"zeroground/colors"
	"zeroground/physics"
	"zeroground/ui"

	"github.com/veandco/go-sdl2/sdl"
)

// firePattern fires around at fixed speed
//	0 1 2
//  3 T 5
//  6 7 8
var firePattern = [8]sdl.Point{
	{X: -1, Y: -1},
	{X: 0, Y: -1},
	{X: 1, Y: -1},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 0, Y: 1},
	{X: -1, Y: 1},
	{X: -1, Y: 0},
}

type Tower struct {
	Size   int32
	Pos    sdl.Point
	Color  sdl.Color
	FireAt int
	Speed  uint64

	lastMoved uint64
	createdAt uint64
}

func (t *Tower) Position() sdl.Point {
	return t.Pos
}

func (t *Tower) Hitbox() []physics.Plane2D {
	fx := t.Pos.X + firePattern[t.FireAt].X*t.Size
	fy := t.Pos.Y + firePattern[t.FireAt].Y*t.Size
	return []physics.Plane2D{
		{Rect: sdl.Rect{
			X: t.Pos.X - t.Size/2,
			Y: t.Pos.Y - t.Size/2,
			W: t.Size,
			H: t.Size,
		}},
		{Rect: sdl.Rect{
			X: fx - t.Size/2,
			Y: fy - t.Size/2,
			W: t.Size,
			H: t.Size,
		}},
	}
}

func (t *Tower) Rect() []sdl.Rect {
	fx := t.Pos.X + firePattern[t.FireAt].X*t.Size
	fy := t.Pos.Y + firePattern[t.FireAt].Y*t.Size
	return []sdl.Rect{
		{
			X: t.Pos.X - t.Size/2,
			Y: t.Pos.Y - t.Size/2,
			W: t.Size,
			H: t.Size,
		},
		{
			X: fx - t.Size/2,
			Y: fy - t.Size/2,
			W: t.Size,
			H: t.Size,
		},
	}
}

func (t *Tower) HasCollision(rects ...sdl.Rect) bool {
	mine := t.Rect()
	for _, r1 := range rects {
		for _, r2 := range mine {
			if r1.HasIntersection(&r2) {
				return true
			}
		}
	}
	return false
}

func (t *Tower) fireAtNextPos() {
	t.FireAt = (t.FireAt + 1) % len(firePattern)
}

func (t *Tower) Handle() {}

func (t *Tower) Update() {
	now := sdl.GetTicks64()
	if now-t.lastMoved >= t.Speed {
		t.fireAtNextPos()
		t.lastMoved = now
	}
}

func (t *Tower) Draw(renderer *sdl.Renderer) {
	rects := t.Rect()
	white := colors.White()
	dark := colors.Darker(t.Color)
	// turret
	renderer.SetDrawColor(colors.RGBA(dark))
	renderer.FillRect(&rects[0])
	// border
	// w := colors.White()
	// w.
	renderer.SetDrawColor(colors.RGBA(white))
	renderer.DrawRect(&rects[0])
	// fire
	renderer.SetDrawColor(colors.RGBA(t.Color))
	renderer.FillRect(&rects[1])
}

func (t *Tower) Reset() {
	t.Pos = sdl.Point{
		X: rand.Int31() % (ui.WindowWidth - 2*t.Size),
		Y: rand.Int31() % (ui.WindowHeight - 2*t.Size),
	}
	t.FireAt = 1
	t.createdAt = sdl.GetTicks64()
	t.lastMoved = sdl.GetTicks64()
}

func NewTower(size int32, color sdl.Color, speed uint64) *Tower {
	t := &Tower{
		Color: color,
		Size:  size,
		Speed: speed,
	}
	t.Reset()
	return t
}

package physics

import (
	"zeroground/pkg/colors"

	"github.com/veandco/go-sdl2/sdl"
)

type Plane2D struct {
	sdl.Rect
}

func (p Plane2D) Draw(r *sdl.Renderer) {
	r.SetDrawColor(colors.RGBA(colors.New(148, 210, 189)))
	// r.SetDrawColor(0, 18, 25, 0)
	r.DrawRect(&p.Rect)
}

func (p Plane2D) BoundingRect() *sdl.Rect {
	return &p.Rect
}

func (p *Plane2D) Center() sdl.Point {
	return sdl.Point{
		X: p.X,
		Y: p.Y,
	}
}

// func (p *Plane2D) Hitbox()

func NewPlane2D(rect sdl.Rect) *Plane2D {
	return &Plane2D{rect}
}

// Hitboxer might be neened if you want to utilise hitboxes in a non-physics context
// type Hitboxer interface {
// 	Hitbox() []Plane2D
// }

type PhysicsObject2D interface {
	Hitbox() []Plane2D
}

func HasIntersection(o1, o2 PhysicsObject2D) bool {
	others := o2.Hitbox()
	for _, h1 := range o1.Hitbox() {
		for _, h2 := range others {
			if h1.Rect.HasIntersection(&h2.Rect) {
				return true
			}
		}
	}
	return false
}

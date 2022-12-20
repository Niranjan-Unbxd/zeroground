package colors

import (
	"github.com/veandco/go-sdl2/sdl"
)

// type Color color.RGBA

// // interface {
// // 	RGBA() (uint8, uint8, uint8, uint8)
// // 	Uint32() uint32
// // }

// // type color color.RGBA

// func (c Color) RGBA() (r, g, b, a uint8) {
// 	return c.r, c.g, c
// }

// // Uint32 return uint32 representation of RGBA color.
// // Borrowed from "sdl.Color"
// func (c Color) Uint32() uint32 {
// 	var v uint32
// 	v |= uint32(c.R) << 24
// 	v |= uint32(c.G) << 16
// 	v |= uint32(c.B) << 8
// 	v |= uint32(c.A)
// 	return v
// }

func New(r, g, b uint8) sdl.Color {
	return sdl.Color{R: r, G: g, B: b, A: sdl.ALPHA_TRANSPARENT}
}

func RGBA(c sdl.Color) (uint8, uint8, uint8, uint8) {
	return c.R, c.G, c.B, c.A
}

func Darker(c sdl.Color) sdl.Color {
	return New(
		uint8(float32(c.R)*1.25)%255,
		uint8(float32(c.G)*1.25)%255,
		uint8(float32(c.B)*1.25)%255,
	)
}

func Red() sdl.Color {
	return New(255, 0, 0)
}

func Green() sdl.Color {
	return New(0, 255, 0)
}

func Blue() sdl.Color {
	return New(0, 0, 255)
}

func Black() sdl.Color {
	return New(0, 0, 0)
}

func White() sdl.Color {
	return New(255, 255, 255)
}

func Grey() sdl.Color {
	return New(189, 189, 189)
}

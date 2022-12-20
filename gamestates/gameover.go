package gamestates

import (
	"zeroground/colors"
	"zeroground/pkg/game"
	"zeroground/ui"

	"github.com/veandco/go-sdl2/sdl"
)

type gameOver struct {
	// font *ttf.Font
	// tex  *sdl.Texture
	// dest sdl.Rect

	menu []ui.Basic

	nextState string
	// renderer  *sdl.Renderer
}

func (g *gameOver) Load() {
	// f, err := ttf.OpenFont("resources/mechfont.otf", 32)
	// if err != nil {
	// 	panic(err)
	// }

	// tex, w, h := game.TextureFromText(g.renderer, g.font, colors.White(), "Game Over!\nR - Restart\nEsc - Exit")

	// g.tex = tex
	// g.dest = sdl.Rect{X: 100, Y: 100, W: w, H: h}
}

func (g *gameOver) Update() string {
	return g.nextState
}

func (g *gameOver) Handle(event sdl.Event) {
	if event.GetType() == sdl.KEYUP {
		ke, _ := event.(*sdl.KeyboardEvent)
		switch ke.Keysym.Scancode {
		case sdl.SCANCODE_R:
			g.nextState = "ingame"
		case sdl.SCANCODE_M:
			g.nextState = "mainmenu"
		}
	}
}

func (g *gameOver) Render(r *sdl.Renderer) {
	r.Clear()
	for _, mtext := range g.menu {
		mtext.Render(r)
	}
}

func (g *gameOver) Free() {
}

func GameStateGameOver(g *game.Game) game.StateFunc {
	return func() game.State {
		white := colors.White()
		return &gameOver{
			menu: []ui.Basic{
				ui.NewText(130, 50, "  Game Over!  ", white, g.Renderer(), g.Font()),
				ui.NewText(130, 200, "  [R] Restart ", white, g.Renderer(), g.Font()),
				ui.NewText(130, 250, " [M] Main Menu", white, g.Renderer(), g.Font()),
				ui.NewText(130, 300, "  [Esc] Exit  ", white, g.Renderer(), g.Font()),
			},
		}
	}
}

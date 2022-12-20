package gamestates

import (
	"zeroground/pkg/colors"
	"zeroground/pkg/game"
	"zeroground/pkg/ui"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type mainMenu struct {
	font *ttf.Font
	tex  *sdl.Texture
	dest sdl.Rect

	menu []ui.Basic

	nextState string
	renderer  *sdl.Renderer
}

func (m *mainMenu) Load() {
	s, err := img.Load("resources/start.png")
	if err != nil {
		panic(err)
	}
	defer s.Free()

	tex, err := m.renderer.CreateTextureFromSurface(s)
	if err != nil {
		panic(err)
	}

	m.tex = tex
	m.dest = sdl.Rect{X: 0, Y: 0, W: s.W, H: s.H}
}

func (m *mainMenu) Handle(e sdl.Event) {
	if e.GetType() == sdl.KEYUP {
		kbe, _ := e.(*sdl.KeyboardEvent)
		switch kbe.Keysym.Scancode {
		case sdl.SCANCODE_S:
			m.nextState = "ingame"
		}
	}
}

func (m *mainMenu) Update() string {
	return m.nextState
}

func (m *mainMenu) Render(r *sdl.Renderer) {
	outline := sdl.Rect{X: m.dest.X + 10, Y: m.dest.Y + 10, W: m.dest.W - 10, H: m.dest.H - 10}

	r.Clear()
	r.Copy(m.tex, nil, &m.dest)
	r.SetDrawColor(255, 255, 255, 0)
	r.DrawRect(&outline)

	for _, mtext := range m.menu {
		mtext.Render(r)
	}
}

func (m *mainMenu) Free() {
	m.tex.Destroy()
}

func GameStateMainMenu(g *game.Game) game.StateFunc {
	return func() game.State {
		// ms := "[S] Start\n[Esc] Exit"

		m := &mainMenu{
			font:     g.Font(),
			renderer: g.Renderer(),
			menu: []ui.Basic{
				ui.NewText(180, 230, " [S] Start", colors.White(), g.Renderer(), g.Font()),
				ui.NewText(180, 280, "[Esc] Exit", colors.White(), g.Renderer(), g.Font()),
			},
		}

		return m
	}
}

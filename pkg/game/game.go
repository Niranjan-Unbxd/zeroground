package game

import (
	"fmt"
	"os"
	"zeroground/colors"
	"zeroground/ui"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type State interface {
	Load()
	Handle(sdl.Event)
	Update() string
	Render(*sdl.Renderer)
	Free()
}

type StateFunc func() State

type Game struct {
	font     *ttf.Font
	renderer *sdl.Renderer
	window   *sdl.Window

	event *sdl.Event

	state State // current game state

	// debugMap map[string]interface{}
	stateMap map[string]StateFunc // map of states

	fps       uint64
	frameRate uint64
	elapsed   uint64
	running   bool
	debug     bool
}

func (g *Game) Renderer() *sdl.Renderer {
	return g.renderer
}

func (g *Game) Font() *ttf.Font {
	return g.font
}

func (g *Game) stop() {
	g.font.Close()
	g.renderer.Destroy()
	g.window.Destroy()
	img.Quit()
	ttf.Quit()
	sdl.Quit()
	// fmt.Println("sdl stop")
}

func (g *Game) handle(event sdl.Event) bool {
	handled := false
	if event.GetType() == sdl.QUIT {
		g.running = false
		handled = true
	}

	if event.GetType() == sdl.KEYUP {
		e, _ := event.(*sdl.KeyboardEvent)
		if e.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
			g.running = false
			handled = true
		}
	}
	return handled
}

func (g *Game) RegisterState(name string, fn StateFunc) {
	g.stateMap[name] = fn
}

func (g *Game) SetState(name string) {
	next, ok := g.stateMap[name]
	if !ok {
		return
	}

	if g.state != nil {
		g.state.Free()
	}
	g.state = next()
	g.state.Load()
}

func (g *Game) Run() {
	var next string

	for g.running {
		start := sdl.GetTicks64()
		for {
			e := sdl.PollEvent()
			if e == nil || g.handle(e) {
				break
			}
			g.state.Handle(e)
		}

		next = g.state.Update()
		if next != "" {
			g.SetState(next)
		}

		g.renderer.SetDrawColor(0, 0, 0, 0)
		g.state.Render(g.renderer)

		if g.debug {
			tex, w, h := TextureFromText(g.renderer, g.font, colors.White(), fmt.Sprintf("FPS: %d", g.fps))
			if tex != nil {
				g.renderer.Copy(tex, nil, &sdl.Rect{X: 10, Y: 10, W: w, H: h})
			}
		}

		g.renderer.Present()

		elapsed := sdl.GetTicks64() - start
		if elapsed < 1 {
			elapsed = 1
		}
		g.fps = 1000 / elapsed
		if g.frameRate > 0 && g.frameRate < g.fps {
			sleep := uint32(1000/g.frameRate) - uint32(elapsed)
			sdl.Delay(sleep)
			elapsed := sdl.GetTicks64() - start
			g.fps = 1000 / elapsed
		}
	}

	g.state.Free()
	g.stop()
}

type GameOption func(*Game) error

// func WithGameState(name string, fn StateFunc) GameOption {
// 	return func(g *Game) error {
// 		g.stateMap[name] = fn
// 		return nil
// 	}
// }

func WithDebug(debug bool) GameOption {
	return func(g *Game) error {
		g.debug = debug
		return nil
	}
}

func WithDefaultFont(path string, size int) GameOption {
	return func(g *Game) error {
		f, err := ttf.OpenFont(path, size)
		if err != nil {
			return err
		}
		g.font = f
		return nil
	}
}

func NewGame(opts ...GameOption) (*Game, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "failed to init SDL2")
		os.Exit(1)
	}

	if err := ttf.Init(); err != nil {
		return nil, err
	}

	if err := img.Init(img.INIT_PNG); err != nil {
		return nil, err
	}

	w, err := sdl.CreateWindow("zeroground v0.1", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, ui.WindowWidth, ui.WindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "failed to create Window")
		return nil, err
	}

	r, err := sdl.CreateRenderer(w, 0, sdl.RENDERER_ACCELERATED)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "failed to create Renderer")
		return nil, err
	}

	// fmt.Println("sdl init")

	g := &Game{
		window:    w,
		renderer:  r,
		event:     nil,
		frameRate: 60,
		running:   true,
		stateMap:  make(map[string]StateFunc),
	}

	for _, opt := range opts {
		err = opt(g)
		if err != nil {
			sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, fmt.Sprintf("failed to apply game options: %v", opt))
			return nil, err
		}
	}

	return g, nil
}

package gamestates

import (
	"zeroground/internal/objects"

	"zeroground/pkg/colors"
	"zeroground/pkg/game"
	"zeroground/pkg/physics"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	boulderSize = int32(32)
)

type inGame struct {
	player  objects.Player
	foods   *objects.FoodSpawner
	tower   *objects.Tower
	spawner *objects.Spawner

	snakeTexture   *sdl.Texture
	boulderTexture *sdl.Texture
	boulders       []*sdl.Rect

	paused bool
}

func (g *inGame) Load() {
	enemyColor := colors.New(187, 62, 3)

	g.player = objects.NewPlayer(sdl.Point{X: 100, Y: 100}, colors.New(10, 147, 150), g.snakeTexture)
	g.foods = objects.NewFoodSpawner()
	g.tower = objects.NewTower(16, enemyColor, 100)
	g.spawner = objects.NewSpawner(enemyColor)
}

func (g *inGame) Free() {}

func (g *inGame) spawnBoulder(x, y int32) {
	rect := &sdl.Rect{
		X: x - boulderSize/2,
		Y: y - boulderSize/2,
		W: boulderSize,
		H: boulderSize,
	}
	g.boulders = append(g.boulders, rect)
}

// func (g *ingame) Reset() {
// 	g.player.Reset()
// 	g.foods.Reset()
// 	g.tower.Reset()
// 	g.spawner.Reset()
// }

func (g *inGame) Handle(event sdl.Event) {
	if event.GetType() == sdl.KEYUP {
		kbe, _ := event.(*sdl.KeyboardEvent)
		if kbe.Keysym.Scancode == sdl.SCANCODE_P {
			g.paused = !g.paused
		}
	}
	if g.paused {
		return
	}
	g.player.Handle(event)
}

func (g *inGame) isGameOver() bool {
	food := g.foods.IntersectingFood(g.player)
	if food != nil {
		g.player.Eat(food)
	}

	if physics.HasIntersection(g.player, g.tower) {
		// fmt.Println("plater hits tower")
		return true
	}

	if physics.HasIntersection(g.player, g.spawner) {
		// fmt.Println("plater hits spawner")
		return true
	}

	_, hit := g.player.HasIntersection(g.boulders...)
	if hit {
		return true
	}

	return !g.player.IsAlive()
}

func (g *inGame) Update() string {
	if g.paused {
		return ""
	}

	nextState := ""

	g.player.Update()
	g.foods.Update()
	g.tower.Update()
	g.spawner.Update()

	if g.isGameOver() {
		nextState = "gameover"
	}

	return nextState
}

func (g *inGame) Render(r *sdl.Renderer) {
	r.Clear()
	g.player.Draw(r)
	g.foods.Draw(r)
	g.tower.Draw(r)
	g.spawner.Draw(r)

	for _, b := range g.boulders {
		r.CopyEx(g.boulderTexture, nil, b, 0, nil, sdl.FLIP_NONE)
		// r.SetDrawColor(255, 255, 255, 0)
		// r.DrawRect(b)
	}
}

func GameStateInGame(rm *game.ResourceManager) game.StateFunc {
	return func() game.State {
		t, _ := rm.GetTexture("block")
		st, _ := rm.GetTexture("snake")
		ing := &inGame{
			snakeTexture:   st.Texture,
			boulderTexture: t.Texture,
			paused:         false,
		}
		ing.spawnBoulder(100, 200)
		return ing
	}
}

package gamestates

import (
	"zeroground/internal/objects"

	"zeroground/pkg/colors"
	"zeroground/pkg/game"
	"zeroground/pkg/physics"

	"github.com/veandco/go-sdl2/sdl"
)

type inGame struct {
	player  objects.Player
	foods   *objects.FoodSpawner
	tower   *objects.Tower
	spawner *objects.Spawner
}

func (g *inGame) Load() {
	g.player = objects.NewPlayer(sdl.Point{X: 100, Y: 100}, colors.Blue())
	g.foods = objects.NewFoodSpawner()
	g.tower = objects.NewTower(16, colors.Red(), 100)
	g.spawner = objects.NewSpawner(colors.Grey())
}

func (g *inGame) Free() {
}

// func (g *ingame) Reset() {
// 	g.player.Reset()
// 	g.foods.Reset()
// 	g.tower.Reset()
// 	g.spawner.Reset()
// }

func (g *inGame) Handle(event sdl.Event) {
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
	return false
}

func (g *inGame) Update() string {
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
}

func GameStateInGame(g *game.Game) game.StateFunc {
	return func() game.State {
		return &inGame{}
	}
}

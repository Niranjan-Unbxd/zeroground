package gamestates

import (
	"fmt"
	"zeroground/colors"
	"zeroground/objects"
	"zeroground/physics"
	"zeroground/pkg/game"

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

func (g *inGame) Update() string {
	g.player.Update()
	g.foods.Update()
	g.tower.Update()
	g.spawner.Update()

	food := g.foods.IntersectingFood(g.player)
	if food != nil {
		g.player.Eat(food)
	}

	if physics.HasIntersection(g.player, g.tower) {
		// g.player.Reset()
		// g.tower.Reset()
		fmt.Println("plater hits tower")
		return "gameover"
	}

	if physics.HasIntersection(g.player, g.spawner) {
		// g.player.Reset()
		// g.spawner.Reset()
		fmt.Println("plater hits spawner")
		return "gameover"
	}

	return ""
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

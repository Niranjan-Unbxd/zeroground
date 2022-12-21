package main

import (
	"zeroground/internal/gamestates"
	"zeroground/pkg/game"
)

const FPS = 1000 / 60

func main() {
	example, err := game.NewGame(
		game.WithDebug(true),
		game.WithDefaultFont("resources/mechfont.otf", 32),
	)
	if err != nil {
		panic(err)
	}

	rm := game.NewResourceManager(example.Renderer())
	rm.LoadTexture("block", "resources/block.png")
	rm.LoadTexture("snake", "resources/snakehead.png")

	example.RegisterState("mainmenu", gamestates.GameStateMainMenu(example))
	example.RegisterState("ingame", gamestates.GameStateInGame(rm))
	example.RegisterState("gameover", gamestates.GameStateGameOver(example))

	example.SetState("mainmenu")
	example.Run()

	rm.UnloadAll()
}

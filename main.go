package main

import (
	"zeroground/internal/gamestates"
	"zeroground/pkg/game"
)

const FPS = 1000 / 60

func main() {
	game, err := game.NewGame(
		game.WithDebug(true),
		game.WithDefaultFont("resources/mechfont.otf", 32),
	)
	if err != nil {
		panic(err)
	}

	game.RegisterState("mainmenu", gamestates.GameStateMainMenu(game))
	game.RegisterState("ingame", gamestates.GameStateInGame(game))
	game.RegisterState("gameover", gamestates.GameStateGameOver(game))

	game.SetState("mainmenu")

	game.Run()
}

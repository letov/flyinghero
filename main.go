package main

import (
	"log"

	"flyinghero/browser"
	"flyinghero/config"
	"flyinghero/game"
	"flyinghero/gameinit"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.Load,
			browser.NewBrowser,
			game.NewGame,
			gameinit.NewGameInit,
		),
		fx.Invoke(startGame),
	).Run()
}

func startGame(g *game.Game, gameInit *gameinit.GameInit, browser *browser.Browser) {
	if err := g.Start(gameInit); err != nil {
		log.Fatalf("Ошибка запуска игры: %v", err)
	}
}

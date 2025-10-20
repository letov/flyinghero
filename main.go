package main

import (
	"log"

	"flyinghero/browser"
	"flyinghero/config"
	"flyinghero/game"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.Load,
			browser.NewBrowser,
			game.NewGame,
		),
		fx.Invoke(startGame),
	).Run()
}

func startGame(g *game.Game) {
	config := g.GetConfig()
	log.Printf("Конфигурация загружена:")
	log.Printf("  URL: %s", config.URL)
	log.Printf("  Интервал: %v", config.Interval)
	log.Printf("  NES клавиши:")
	log.Printf("    Движение: %v", config.NESKeys.GetMovementKeys())
	log.Printf("    Действия: %v", config.NESKeys.GetActionKeys())
	log.Printf("    Система: %v", config.NESKeys.GetSystemKeys())

	if err := g.Start(); err != nil {
		log.Fatalf("Ошибка запуска игры: %v", err)
	}
}

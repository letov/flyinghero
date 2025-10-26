package game

import (
	"log"
	"time"

	"flyinghero/browser"
	"flyinghero/config"
	"flyinghero/gameinit"
)

type Game struct {
	browser *browser.Browser
	config  *config.Config
}

func NewGame(browser *browser.Browser, config *config.Config) *Game {
	return &Game{
		browser: browser,
		config:  config,
	}
}

func (g *Game) Start(gameInit *gameinit.GameInit) error {
	if err := g.browser.OpenURL(g.config.URL); err != nil {
		return err
	}

	ctx := g.browser.GetContext()

	if err := gameInit.Init(ctx, g.config.GameElement); err != nil {
		log.Printf("Ошибка инициализации игры: %v", err)
	}

	for !gameInit.WaitUntil(ctx) {
		log.Println("Ожидаем завершения инициализации...")
	}

	log.Println("Инициализация завершена, запускаем игровой цикл")
	g.runGameLoop()
	return nil
}

func (g *Game) GetConfig() *config.Config {
	return g.config
}

func (g *Game) runGameLoop() {
	log.Println("Запускаем основной цикл игры...")

	ticker := time.NewTicker(g.config.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			nesKeys := g.config.NESKeys.GetMovementKeys()
			if err := g.browser.PressKeysToElement(nesKeys, g.config.GameElement); err != nil {
				log.Printf("Ошибка нажатия клавиш движения: %v", err)
			}
		}
	}
}

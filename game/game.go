package game

import (
	"log"
	"time"

	"flyinghero/browser"
	"flyinghero/config"
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

func (g *Game) Start() error {
	if err := g.browser.OpenURL(g.config.URL); err != nil {
		return err
	}

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
			if err := g.browser.TakeScreenshot(); err != nil {
				log.Printf("Ошибка создания скриншота: %v", err)
			}

			nesKeys := g.config.NESKeys.GetMovementKeys()
			if err := g.browser.PressKeys(nesKeys); err != nil {
				log.Printf("Ошибка нажатия клавиш: %v", err)
			}
		}
	}
}

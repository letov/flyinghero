package gameinit

import (
	"context"
	"log"
	"time"

	"flyinghero/browser"

	"github.com/chromedp/chromedp"
)

type GameInit struct {
	browser *browser.Browser
}

func NewGameInit(browser *browser.Browser) *GameInit {
	return &GameInit{browser: browser}
}

func (g *GameInit) Init(ctx context.Context, selector string) error {
	log.Println("Ждем загрузки страницы...")

	err := chromedp.Run(ctx,
		chromedp.WaitVisible("body", chromedp.ByQuery),
	)

	if err != nil {
		log.Printf("Ошибка ожидания загрузки страницы: %v", err)
		return err
	}

	time.Sleep(2000 * time.Millisecond)

	log.Printf("Нажимаем Start 2 раза в %s...", selector)

	if err := g.browser.PressKeysToElement([]string{"Enter"}, selector); err != nil {
		log.Printf("Ошибка нажатия Enter: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	if err := g.browser.PressKeysToElement([]string{"Enter"}, selector); err != nil {
		log.Printf("Ошибка нажатия Enter: %v", err)
	}

	return nil
}

func (g *GameInit) WaitUntil(ctx context.Context) bool {
	log.Println("Ожидаем 5 секунд...")
	time.Sleep(5 * time.Second)
	return true
}

package browser

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

type Browser struct {
	ctx context.Context
}

func NewBrowser() *Browser {
	ctx, _ := chromedp.NewContext(context.Background())
	return &Browser{ctx: ctx}
}

func (b *Browser) OpenURL(url string) error {
	log.Printf("Открываем браузер и переходим на %s", url)

	return chromedp.Run(b.ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
	)
}

func (b *Browser) TakeScreenshot() error {
	var buf []byte

	err := chromedp.Run(b.ctx,
		chromedp.CaptureScreenshot(&buf),
	)

	if err != nil {
		return err
	}

	if err := os.MkdirAll("screenshots", 0755); err != nil {
		return err
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("screenshots/screenshot_%s.png", timestamp)

	if err := os.WriteFile(filename, buf, 0644); err != nil {
		return err
	}

	log.Printf("Скриншот сохранен: %s", filename)
	return nil
}

func (b *Browser) PressKeys(keys []string) error {
	for _, key := range keys {
		log.Printf("Нажимаем клавишу: %s", key)

		err := chromedp.Run(b.ctx,
			chromedp.KeyEvent(key),
		)

		if err != nil {
			return err
		}

		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func (b *Browser) Close() {
	chromedp.Cancel(b.ctx)
}

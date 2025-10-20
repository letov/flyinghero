package browser

import (
	"context"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/chromedp/chromedp"
)

type Browser struct {
	ctx            context.Context
	lastScreenshot string
	viewerProcess  *exec.Cmd
}

func NewBrowser() *Browser {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("start-fullscreen", true),
		chromedp.Flag("kiosk", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	return &Browser{ctx: ctx, lastScreenshot: "", viewerProcess: nil}
}

func (b *Browser) OpenURL(url string) error {
	log.Printf("Открываем браузер и переходим на %s", url)

	err := chromedp.Run(b.ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
	)

	if err != nil {
		log.Printf("Ошибка открытия URL: %v", err)
		return err
	}

	log.Printf("Браузер успешно открыт")
	return nil
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

	filename := "screenshots/current.png"

	if err := os.WriteFile(filename, buf, 0644); err != nil {
		return err
	}

	b.lastScreenshot = filename
	log.Printf("Скриншот обновлен: %s", filename)

	b.openScreenshot(filename)
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

func (b *Browser) openScreenshot(filename string) {
	if b.viewerProcess == nil && runtime.GOOS == "darwin" {
		b.viewerProcess = exec.Command("open", "-W", filename)
		go b.viewerProcess.Run()
	}
}

func (b *Browser) GetLastScreenshot() string {
	return b.lastScreenshot
}

func (b *Browser) Close() {
	if b.viewerProcess != nil {
		b.viewerProcess.Process.Kill()
	}
	chromedp.Cancel(b.ctx)
}

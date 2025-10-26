package browser

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

type Browser struct {
	ctx context.Context
}

func NewBrowser() *Browser {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("start-fullscreen", true),
		chromedp.Flag("auto-open-devtools-for-tabs", true),
		chromedp.Flag("devtools", true),
		chromedp.Flag("remote-debugging-port", "9222"),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	return &Browser{ctx: ctx}
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

func (b *Browser) LogToBrowser(message string) {
	log.Printf(message)
	jsCode := `console.log("` + message + `");`
	chromedp.Run(b.ctx, chromedp.Evaluate(jsCode, nil))
}

func (b *Browser) PressKeysToElement(keys []string, selector string) error {
	for _, key := range keys {
		keyCode := b.getKeyCode(key)

		jsCode := `
			document.querySelector("` + selector + `").dispatchEvent(new KeyboardEvent('keydown',{key:'` + key + `',keyCode:` + keyCode + `,bubbles:true}));setTimeout(()=>document.querySelector("` + selector + `").dispatchEvent(new KeyboardEvent('keyup',{key:'` + key + `',keyCode:` + keyCode + `,bubbles:true})),50);
		`

		err := chromedp.Run(b.ctx, chromedp.Evaluate(jsCode, nil))
		if err != nil {
			b.LogToBrowser("Ошибка отправки клавиши " + key + ": " + err.Error())
		}

		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func (b *Browser) getKeyCode(key string) string {
	keyCodeMap := map[string]string{
		"ArrowUp":    "38",
		"ArrowDown":  "40",
		"ArrowLeft":  "37",
		"ArrowRight": "39",
		"Enter":      "13",
		"Space":      "32",
		"KeyZ":       "90",
		"KeyX":       "88",
	}

	if code, exists := keyCodeMap[key]; exists {
		return code
	}
	return "0"
}

func (b *Browser) GetContext() context.Context {
	return b.ctx
}

func (b *Browser) Close() {
	chromedp.Cancel(b.ctx)
}

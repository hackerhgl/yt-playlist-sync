package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

const PLAYLIST = "PLoSjAzdJQCyfsUatRxBnaes_4yz3-ONm-"

func main()  {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag(`headless`, false),
        chromedp.DisableGPU,
        chromedp.Flag(`disable-extensions`, false),
        chromedp.Flag(`enable-automation`, false),
    )


	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()
	
	ctx, cancel = context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	var text string
	input := `//input[@name="search_string"]`

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://ytplaylist-len.herokuapp.com/`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(input),
		// find and click "Example" link
		chromedp.Click(input, chromedp.NodeVisible),
		// retrieve the text of the textarea
		chromedp.SendKeys(input, PLAYLIST),
		chromedp.Submit(input),
		// chromedp.WaitNotPresent(`//*div[@class="mt-4"]//div[@class="container"]//p`,chromedp.BySearch),
		chromedp.WaitVisible(`body .mt-4 .container > p`),
		// chromedp.WaitVisible(`//*div[@class="mt-4"]//div[@class="container"]//p[contains(., 'Search more than')]`),
		chromedp.Text(`body .mt-4 .container > p`, &text),
	)

	fmt.Println("VIDEOS", text)
	if err != nil {
		log.Fatal(err)
	}
}
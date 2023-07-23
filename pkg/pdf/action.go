package pdf

import (
	"context"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func makePdf(options Options, res *[]byte) chromedp.Tasks {
	var actions []chromedp.Action

	if options.Viewport != nil {
		actions = append(actions, chromedp.EmulateViewport(
			int64(options.Viewport.Width),
			int64(options.Viewport.Height),
		))
	}

	actions = append(actions, chromedp.Navigate(options.Url))

	if options.Wait > 0 {
		actions = append(actions, chromedp.Sleep(options.Wait))
	}

	actions = append(actions, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		*res, _, err = page.PrintToPDF().WithPrintBackground(false).Do(ctx)
		if err != nil {
			return err
		}
		return nil
	}))

	return actions
}

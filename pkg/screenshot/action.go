package screenshot

import (
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

func makeScreenshot(options Options, res *[]byte) chromedp.Tasks {
	var actions []chromedp.Action

	if options.Viewport != nil {
		actions = append(actions, chromedp.EmulateViewport(
			int64(options.Viewport.Width),
			int64(options.Viewport.Height),
		))
	}

	if options.DarkMode {
		actions = append(actions, emulation.SetAutoDarkModeOverride().WithEnabled(options.DarkMode))
	}

	actions = append(actions, chromedp.Navigate(options.Url))
	actions = append(actions, chromedp.CaptureScreenshot(res))

	return actions
}

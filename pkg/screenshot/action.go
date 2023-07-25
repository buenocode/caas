package screenshot

import (
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/security"
	"github.com/chromedp/chromedp"
)

func MakeScreenshot(options Options, res *[]byte) chromedp.Tasks {
	var actions []chromedp.Action

	if options.Viewport != nil {
		actions = append(actions, chromedp.EmulateViewport(
			int64(options.Viewport.Width),
			int64(options.Viewport.Height),
		))
	}

	if options.DarkMode {
		actions = append(actions, emulation.SetAutoDarkModeOverride().WithEnabled(true))
	}

	if options.Insecure {
		actions = append(actions, security.SetIgnoreCertificateErrors(true))
	}

	if !options.JavaScript {
		actions = append(actions, emulation.SetScriptExecutionDisabled(true))
	}

	actions = append(actions, chromedp.Navigate(options.Url))

	if options.Wait > 0 {
		actions = append(actions, chromedp.Sleep(options.Wait))
	}

	if options.Full {
		actions = append(actions, chromedp.FullScreenshot(res, 100))
	} else if options.Selector != "" {
		actions = append(actions, chromedp.Screenshot(options.Selector, res, chromedp.ByQuery))
	} else {
		actions = append(actions, chromedp.CaptureScreenshot(res))
	}

	return actions
}

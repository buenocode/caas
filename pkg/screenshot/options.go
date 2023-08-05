package screenshot

import (
	"time"

	"github.com/chromedp/cdproto/network"
)

type Viewport struct {
	Width  int
	Height int
}

type Options struct {
	Url        string
	Viewport   *Viewport
	Cookies    []network.SetCookieParams
	DarkMode   bool
	Full       bool
	Selector   string
	Insecure   bool
	JavaScript bool
	Wait       time.Duration
}

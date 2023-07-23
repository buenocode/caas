package screenshot

import "time"

type Viewport struct {
	Width  int
	Height int
}

type Options struct {
	Url      string
	Viewport *Viewport
	DarkMode bool
	Wait     time.Duration
}

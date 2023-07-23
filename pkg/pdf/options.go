package pdf

import "time"

type Viewport struct {
	Width  int
	Height int
}

type Options struct {
	Url      string
	Viewport *Viewport
	Wait     time.Duration
}

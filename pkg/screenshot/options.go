package screenshot

import "time"

type Viewport struct {
	Width  int
	Height int
}

type Options struct {
	Url        string
	Viewport   *Viewport
	DarkMode   bool
	Full       bool
	Selector   string
	Insecure   bool
	JavaScript bool
	Wait       time.Duration
}

package screenshot

type Viewport struct {
	Width  int
	Height int
}

type Options struct {
	Url      string
	Viewport *Viewport
	DarkMode bool
}

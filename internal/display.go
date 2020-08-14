package internal

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	defaultWidth = 600
	defaultHeight = 800
)

// SDLDisplay struct that represents an SDL display
type SDLDisplay struct{
	Window *sdl.Window
	Renderer *sdl.Renderer
}

// Destroy destroys the SDLDisplay and it's components
func (d *SDLDisplay) Destroy() {
	if d.Window != nil {
		defer d.Window.Destroy()
	}

	if d.Renderer != nil {
		defer d.Renderer.Destroy()
	}
}

// MakeDisplay creates a new SDLDisplay
func MakeDisplay(flags uint32) (*SDLDisplay, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	window, err := sdl.CreateWindow(
		"Gaming in go",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		defaultHeight,
		defaultHeight,
		flags,
	)

	if err != nil {
		return nil, err
	}


	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		return nil, err
	}

	return &SDLDisplay{
		Window: window,
		Renderer: renderer,
	}, nil
}
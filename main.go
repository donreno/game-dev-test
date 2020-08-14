package main

import (
	"fmt"
	"game-dev-test/internal"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	display, err := internal.MakeDisplay(sdl.WINDOW_OPENGL)
	defer func() {
		if display != nil{
			display.Destroy()
		}
	}()

	if err != nil{
		fmt.Printf("SDL Init error:%v\n", err)
		return
	}
	
	
	player, err := internal.MakePlayer(display, "sprites/player.bmp")
	if err != nil {
		fmt.Printf("Error loading assets:%v\n", err)
	}
	
	defer player.Destroy()

	gameloop(display, player)
}

func gameloop(d *internal.SDLDisplay, player *internal.Player) {
	
	player.X = 40
	player.Y = 20

	for{
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quiting... Good bye!")
				return
			}
		}

		d.Renderer.SetDrawColor(255,255,255,255)
		d.Renderer.Clear()

		player.Draw(d.Renderer)

		d.Renderer.Present()
	}
}
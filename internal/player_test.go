package internal_test

import (
	"game-dev-test/internal"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestPlayer(t *testing.T){
	display,_ := internal.MakeDisplay(sdl.WINDOW_HIDDEN)
	defer display.Destroy()

	t.Run("It should create player with no errors", func(t *testing.T){
		p, err := internal.MakePlayer(display, "../sprites/player.bmp")
		defer p.Destroy()
		
		assert.NoError(t, err)
		assert.NotNil(t, p)
	})

	t.Run("It should return an error if sprite path is incorrect", func(t *testing.T){
		p, err := internal.MakePlayer(display, "sprites/fake_file.bmp")
		
		assert.Error(t, err)
		assert.Nil(t, p)
	})

	t.Run("It should draw on renderer", func(t *testing.T){
		p, _ := internal.MakePlayer(display, "../sprites/player.bmp")
		defer p.Destroy()

		p.Draw(display.Renderer)
	})
}
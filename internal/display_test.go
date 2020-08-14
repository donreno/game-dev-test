package internal_test

import (
	"game-dev-test/internal"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
)

func TestDisplay(t *testing.T){
	t.Run("It should create Display properly", func(t *testing.T){
		d, err := internal.MakeDisplay(sdl.WINDOW_HIDDEN)
		defer d.Destroy()

		assert.NoError(t, err)
		assert.NotNil(t, d)
	})
}
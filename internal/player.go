package internal

import "github.com/veandco/go-sdl2/sdl"

// Player struct that represents a player
type Player struct {
	Texture *sdl.Texture
	X int32
	Y int32
}

// Destroy player destroy func
func (p *Player) Destroy(){
	if p.Texture != nil {
		p.Texture.Destroy()
	}
}

// Draw draws a player on the renderer
func (p *Player) Draw(r *sdl.Renderer) error {
	return r.Copy(
		p.Texture,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: p.X, Y: p.Y, W: 105, H: 105},
	)
}

// MakePlayer creates a new player loading it's assets
func MakePlayer(d *SDLDisplay, spritepath string) (*Player, error){
	img, err := sdl.LoadBMP(spritepath)
	if err != nil{
		return nil, err
	}

	defer img.Free()

	texture, err := d.Renderer.CreateTextureFromSurface(img)
	if err != nil{
		return nil, err
	}

	return &Player{
		Texture: texture,
	}, nil
}
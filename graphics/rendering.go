package graphics

import (
	"github.com/veandco/go-sdl2/sdl"
)

var camera = NewCamera()

func RenderScene(renderer *sdl.Renderer) {
	renderer.Clear()
	// TODO: RenderScene here
	renderer.Present()
}

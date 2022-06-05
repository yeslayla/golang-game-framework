package game

import (
	"log"

	"github.com/manleydev/golang-game-framework/node"
	"github.com/manleydev/golang-game-framework/rendering"
	"github.com/manleydev/golang-game-framework/sdl"
)

func Run(root *node.Node, renderer rendering.Renderer2D) {
	texture := sdl.NewSdlTexture2D(renderer, "gopher.bmp")

	g := node.NewSprite2D(texture)
	defer root.AddChild(g)

	g.OnReady(func() error {
		log.Println("Hello world!")
		return nil
	})

	g.OnUpdate(func() error {
		g.Rotation += 0.01
		return nil
	})

	g.Name = "Gopher"

	camera := node.NewCamera2D(&renderer, rendering.NewStandardCamera())
	defer root.AddChild(camera)

	camera.OnReady(func() error {
		return camera.Enable()
	})

}

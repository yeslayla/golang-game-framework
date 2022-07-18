package game

import (
	"log"

	"github.com/manleydev/golang-game-framework/input"
	"github.com/manleydev/golang-game-framework/node"
	"github.com/manleydev/golang-game-framework/rendering"
	"github.com/manleydev/golang-game-framework/sdl"
)

func Run(root *node.Node, renderer rendering.Renderer2D, inputHandler input.InputHandler) {
	texture := sdl.NewSdlTexture2D(renderer, "gopher.bmp")

	g := node.NewSprite2D(texture)
	g.Name = "Gopher"
	defer root.AddChild(g)

	g.OnReady(func() error {
		log.Println("Hello world!")
		return nil
	})

	speed := 0.01
	g.OnUpdate(func() error {

		if inputHandler.IsKeyJustPressed(sdl.SCANCODE_SPACE) {
			speed = -speed
		}

		g.Rotation += speed
		return nil
	})

	controller := node.NewNode()
	controller.Name = "Controller"
	defer root.AddChild(&controller)

	controller.OnUpdate(func() error {
		if inputHandler.IsKeyJustPressed(sdl.SCANCODE_ESCAPE) {
			if g.IsProcessing() {
				g.SetProcessMode(node.PausedProcessMode)
			} else {
				g.SetProcessMode(node.DefaultProcessMode)
			}
		}

		return nil
	})

	camera := node.NewCamera2D(&renderer, rendering.NewStandardCamera())
	defer root.AddChild(camera)

	camera.OnReady(func() error {
		return camera.Enable()
	})

}

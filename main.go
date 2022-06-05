package main

import (
	"log"

	"github.com/manleydev/golang-game-framework/game"
	"github.com/manleydev/golang-game-framework/node"
	"github.com/manleydev/golang-game-framework/rendering"
	"github.com/manleydev/golang-game-framework/sdl"
)

func main() {
	var renderer rendering.Renderer2D = sdl.NewSdlRenderer2D(sdl.SdlRenderer2DInput{
		WindowTitle:  "Sample Game",
		WindowWidth:  1280,
		WindowHeight: 720,
	})
	defer renderer.Destroy()

	root := node.NewNode()
	root.Name = "Root"

	game.Run(&root, renderer)

	running := true
	for running {
		if err := root.Update(); err != nil {
			log.Fatal("Update: ", err)
		}
		if err := renderer.Update(); err != nil {
			log.Fatal("Renderer Update: ", err)
		}

		if err := root.Draw(renderer); err != nil {
			log.Fatal("Draw: ", err)
		}
		if err := renderer.Draw(); err != nil {
			log.Fatal("Renderer Draw: ", err)
		}

	}
}

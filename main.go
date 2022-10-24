package main

import (
	"log"
	"time"

	"github.com/manleydev/golang-game-framework/game"
	"github.com/manleydev/golang-game-framework/input"
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

	var input input.InputHandler = sdl.NewSdlInputHandler()

	root := node.NewNode()
	root.Name = "Root"

	game.Run(&root, renderer, input)

	running := true
	var lastTimestamp int64 = time.Now().UnixMicro()
	var delta float64 = 0.0
	for running {

		if err := root.Update(delta); err != nil {
			log.Fatal("Update: ", err)
		}
		if err := input.Update(delta); err != nil {
			log.Fatal("Input Update: ", err)
		}

		if err := renderer.Update(delta); err != nil {
			log.Fatal("Renderer Update: ", err)
		}

		if err := root.Draw(renderer); err != nil {
			log.Fatal("Draw: ", err)
		}
		if err := renderer.Draw(); err != nil {
			log.Fatal("Renderer Draw: ", err)
		}

		newTimestamp := time.Now().UnixMicro()
		delta = float64(lastTimestamp) / float64(newTimestamp)
		lastTimestamp = newTimestamp

	}
}

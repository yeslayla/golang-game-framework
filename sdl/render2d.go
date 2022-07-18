package sdl

import (
	"errors"
	"log"

	"github.com/manleydev/golang-game-framework/core"
	"github.com/manleydev/golang-game-framework/rendering"
	"github.com/veandco/go-sdl2/sdl"
)

type SdlRenderer2D struct {
	window        *sdl.Window
	renderer      *sdl.Renderer
	currentCamera *rendering.Camera2D
	inputHandler  *SdlEventHandler
}

type SdlRenderer2DInput struct {
	WindowTitle  string
	WindowWidth  int32
	WindowHeight int32
	Fullscreen   bool
}

func NewSdlRenderer2D(input SdlRenderer2DInput) *SdlRenderer2D {
	w := &SdlRenderer2D{}
	var err error

	w.window, err = sdl.CreateWindow(input.WindowTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		input.WindowWidth, input.WindowHeight,
		sdl.WINDOW_VULKAN)
	if err != nil {
		log.Print("Failed to create SDL window: ", err)
		return nil
	}

	if input.Fullscreen {
		if err := w.window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP); err != nil {
			log.Print("Failed to set fullscreen: ", err)
			return nil
		}
	}

	w.window.SetResizable(true)

	w.renderer, err = sdl.CreateRenderer(w.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Print("Failed to create SDL renderer: ", err)
		return nil
	}

	return w
}

func (w *SdlRenderer2D) Destroy() {
	if err := w.window.Destroy(); err != nil {
		log.Print("Failed to destroy SDL window: ", err)
	}

	if err := w.renderer.Destroy(); err != nil {
		log.Print("Failed to destroy SDL renderer: ", err)
	}
}

func (w *SdlRenderer2D) SetCamera(camera *rendering.Camera2D) error {
	w.currentCamera = camera
	return nil
}

func (w *SdlRenderer2D) Update() error {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return errors.New("Quit system not yet implemented!")
		default:
			if w.inputHandler != nil {
				handler := *w.inputHandler
				handler(event)
			}
		}
	}
	return nil
}

func (w *SdlRenderer2D) DrawTexture2D(input rendering.DrawTexture2DInput) error {
	tex, ok := input.Texture.(*SdlTexture2D)
	if !ok {
		return errors.New("Texture is not an SdlTexture2D!")
	}

	center := tex.GetCenter()
	modifiers := rendering.Camera2DModifiers{
		Zoom:   1.0,
		Offset: core.Vector2{X: 0, Y: 0},
	}
	if w.currentCamera != nil {
		modifiers = (*w.currentCamera).GetModifiers()
	}

	if err := w.renderer.CopyEx(tex.texture,
		&sdl.Rect{
			X: int32(input.Rect.X),
			Y: int32(input.Rect.Y),
			W: int32(input.Rect.W),
			H: int32(input.Rect.H),
		},
		&sdl.Rect{
			X: int32((input.Position.X - modifiers.Offset.X) * modifiers.Zoom),
			Y: int32((input.Position.Y - modifiers.Offset.Y) * modifiers.Zoom),
			W: int32(input.Rect.W * modifiers.Zoom),
			H: int32(input.Rect.H * modifiers.Zoom),
		}, input.Rotation,
		&sdl.Point{
			X: int32(center.X * modifiers.Zoom),
			Y: int32(center.Y * modifiers.Zoom),
		},
		sdl.FLIP_NONE); err != nil {
		return err
	}

	return nil
}

func (w *SdlRenderer2D) Draw() error {

	w.renderer.Present()

	if err := w.renderer.SetDrawColor(0, 0, 0, 255); err != nil {
		return err
	}

	if err := w.renderer.Clear(); err != nil {
		return err
	}

	return nil
}

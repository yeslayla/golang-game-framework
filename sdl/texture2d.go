package sdl

import (
	"io/ioutil"
	"log"

	"github.com/manleydev/golang-game-framework/core"
	"github.com/manleydev/golang-game-framework/importer"
	"github.com/manleydev/golang-game-framework/rendering"

	"github.com/veandco/go-sdl2/sdl"
)

type SdlTexture2D struct {
	texture *sdl.Texture
}

func NewSdlTexture2D(renderer rendering.Renderer2D, bmpPath string) *SdlTexture2D {
	sdlRenderer, ok := renderer.(*SdlRenderer2D)
	if !ok {
		log.Print("Renderer is not an SDL renderer")
		return nil
	}

	tex := SdlTexture2D{}
	surface, err := sdl.LoadBMP(bmpPath)
	if err != nil {
		log.Print("Failed to load bmp: ", err)
		return nil
	}
	defer surface.Free()

	tex.texture, err = sdlRenderer.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Print("Failed to create SDL texture: ", err)
		return nil
	}

	return &tex
}

func NewEmbededSdlTexture2D(renderer rendering.Renderer2D, bmpPath string) *SdlTexture2D {
	sdlRenderer, ok := renderer.(*SdlRenderer2D)
	if !ok {
		log.Print("Renderer is not an SDL renderer")
		return nil
	}

	f, err := importer.EmbededAssets.Open(bmpPath)
	if err != nil {
		log.Printf("Failed to load '%s' from embeded assets: %s", bmpPath, err)
		return nil
	}
	defer f.Close()

	raw, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("Failed to read '%s' from embeded assets: %s", bmpPath, err)
		return nil
	}

	buffer, err := sdl.RWFromMem(raw)
	if err != nil {
		log.Printf("Failed to create buffer for embeded asset '%s': %s", bmpPath, err)
	}

	tex := SdlTexture2D{}
	surface, err := sdl.LoadBMPRW(buffer, true)
	if err != nil {
		log.Print("Failed to load bmp: ", err)
		return nil
	}
	defer surface.Free()

	tex.texture, err = sdlRenderer.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Print("Failed to create SDL texture: ", err)
		return nil
	}

	return &tex
}

func (tex SdlTexture2D) GetCenter() core.Vector2 {
	_, _, width, height, err := tex.texture.Query()
	if err != nil {
		log.Print("Could not get texture center: ", err)
		return core.Vector2{}
	}

	return core.Vector2{
		X: float64(width / 2),
		Y: float64(height / 2),
	}
}

func (tex SdlTexture2D) GetRect() core.Rect2D {
	_, _, width, height, err := tex.texture.Query()
	if err != nil {
		log.Print("Could not get texture data: ", err)
		return core.Rect2D{}
	}
	return core.Rect2D{
		X: 0,
		Y: 0,
		W: float64(width),
		H: float64(height),
	}
}

func (tex SdlTexture2D) Destroy() {
	if err := tex.texture.Destroy(); err != nil {
		log.Print("Failed to destroy SDL texture: ", err)
	}
}

package sdl

import (
	"github.com/veandco/go-sdl2/sdl"
)

const EXPECTED_KEYBOARD_STATE_SIZE = 512

type SdlInputHandler struct {
	keyState     []uint8
	lastKeyState []uint8
}

func (input *SdlInputHandler) IsKeyDown(key uint) bool {
	keys := sdl.GetKeyboardState()

	if keys[key] == 1 {
		return true
	}

	return false
}

func (input *SdlInputHandler) IsKeyJustPressed(key uint) bool {
	if input.keyState[key] == 1 && input.lastKeyState[key] == 0 {
		return true
	}

	return false
}

func (input *SdlInputHandler) IsKeyJustReleased(key uint) bool {
	if input.keyState[key] == 0 && input.lastKeyState[key] == 1 {
		return true
	}

	return false
}

func (input *SdlInputHandler) Update(delta float64) error {
	copy(input.lastKeyState, input.keyState)
	input.keyState = sdl.GetKeyboardState()

	return nil
}

func NewSdlInputHandler() *SdlInputHandler {
	handler := &SdlInputHandler{}
	handler.lastKeyState = make([]uint8, EXPECTED_KEYBOARD_STATE_SIZE)
	handler.keyState = make([]uint8, EXPECTED_KEYBOARD_STATE_SIZE)

	return handler
}

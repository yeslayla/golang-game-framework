package input

type InputHandler interface {
	IsKeyDown(uint) bool
	IsKeyJustPressed(uint) bool
	IsKeyJustReleased(uint) bool
	Update() error
}

package input

type InputHandler interface {
	IsKeyDown(uint32) bool
	IsKeyJustPressed(uint32) bool
	IsKeyJustReleased(uint32) bool
	Update(delta float64) error
}

package core

import "fmt"

type Rect2D struct {
	X float64
	Y float64
	W float64
	H float64
}

func (rect Rect2D) String() string {
	return fmt.Sprintf("X: %v, Y: %v, W: %v, H: %v", rect.X, rect.Y, rect.W, rect.H)
}

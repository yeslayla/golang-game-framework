package core

import "fmt"

type Vector2 struct {
	X float64
	Y float64
}

func (p Vector2) String() string {
	return fmt.Sprintf("X: %v, Y: %v", p.X, p.Y)
}

func (vector Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		X: vector.X + other.X,
		Y: vector.Y + other.Y,
	}
}

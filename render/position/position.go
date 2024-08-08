package position

type Position struct {
	X, Y float32
}

func (p *Position) Move(dx, dy float32) {
	p.X += dx
	p.Y += dy
}

func NewPosition(x, y float32) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

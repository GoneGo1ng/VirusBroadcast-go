package point

// 在图中的位置
type Point struct {
	X int // x轴坐标
	Y int // y轴坐标
}

// 移动到图中另一个位置
func (p *Point) MoveTo(x, y int) {
	p.X += x
	p.Y += y
}


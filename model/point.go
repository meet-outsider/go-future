package model

import "math"

type Point struct{ X, Y float64 }

// Distance 这是给struct Point类型定义一个方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

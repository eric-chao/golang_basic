package gotchas

import (
	"math"
	"testing"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

/*
Since methods with pointer receivers take either a value or a pointer,
you could also skip the pointer altogether:
    var p Point // has zero value Point{X:0, Y:0}
*/
func TestPanic(t *testing.T) {
	var p = new(Point) // var p *Point  will panic
	t.Log(p.Abs())

	var p1 Point
	t.Log(p1.Abs())
}

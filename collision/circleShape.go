package collision

import (
	"github.com/teomat/mater/aabb"
	"github.com/teomat/mater/transform"
	"github.com/teomat/mater/vect"
)

type CircleShape struct {
	Shape *Shape
	// Center of the circle. Call Update() on the parent shape if changed.
	Position vect.Vect
	// Radius of the circle. Call Update() on the parent shape if changed.
	Radius float64
	// Global center of the circle. Do not touch!
	Tc vect.Vect
}

// Creates a new CircleShape with the given center and radius.
func NewCircle(pos vect.Vect, radius float64) *Shape {
	shape := newShape()
	circle := &CircleShape{
		Position: pos,
		Radius:   radius,
		Shape:    shape,
	}
	shape.ShapeClass = circle
	return shape
}

// Returns ShapeType_Circle. Needed to implemet the ShapeClass interface.
func (circle *CircleShape) ShapeType() ShapeType {
	return ShapeType_Circle
}

// Recalculates the global center of the circle and the the bounding box.
func (circle *CircleShape) update(xf transform.Transform) aabb.AABB {
	//global center of the circle
	center := xf.TransformVect(circle.Position)
	circle.Tc = center
	rv := vect.Vect{circle.Radius, circle.Radius}

	return aabb.AABB{
		vect.Sub(center, rv),
		vect.Add(center, rv),
	}
}

// Returns true if the given point is located inside the circle.
func (circle *CircleShape) TestPoint(point vect.Vect) bool {
	d := vect.Sub(point, circle.Tc)

	return vect.Dot(d, d) <= circle.Radius*circle.Radius
}

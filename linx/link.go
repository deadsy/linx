//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------

package linx

//-----------------------------------------------------------------------------

// Body1 is a rigid body pivoting about a point.
type Body1 struct {
	Pivot  Cartesian // The point about which the body pivots.
	Theta  float64   // The angle of the body.
	Points []Polar   // Points on the body relative to the pivot point.
}

// NewBody1 returns a body that pivots about a point.
func NewBody1(pivot Cartesian, points []Cartesian) *Body1 {
	pp := make([]Polar, len(points))
	for i, v := range points {
		v = v.Sub(pivot)
		pp[i] = v.CartesianToPolar()
	}
	return &Body1{
		Pivot:  pivot,
		Points: pp,
	}
}

// Translate body1.
func (b *Body1) Translate(t Cartesian) {
	b.Pivot = b.Pivot.Add(t)
}

// Rotate body1.
func (b *Body1) Rotate(theta float64) {
	b.Theta += theta
}

//-----------------------------------------------------------------------------

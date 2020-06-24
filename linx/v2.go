//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------

package linx

import "math"

//-----------------------------------------------------------------------------

type Polar struct {
	Radius float64
	Theta  float64
}

type Cartesian struct {
	X float64
	Y float64
}

//-----------------------------------------------------------------------------

// PolarToCartesian converts polar to cartesian coordinates.
func (p Polar) PolarToCartesian() Cartesian {
	return Cartesian{
		X: p.Radius * math.Cos(p.Theta),
		Y: p.Radius * math.Sin(p.Theta),
	}
}

// CartesianToPolar converts cartesian to polar coordinates.
func (c Cartesian) CartesianToPolar() Polar {
	return Polar{
		Radius: math.Sqrt(c.X*c.X + c.Y*c.Y),
		Theta:  math.Atan2(c.Y, c.X),
	}
}

//-----------------------------------------------------------------------------

// Add returns a + b
func (a Cartesian) Add(b Cartesian) Cartesian {
	return Cartesian{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

// Sub returns a - b
func (a Cartesian) Sub(b Cartesian) Cartesian {
	return Cartesian{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

//-----------------------------------------------------------------------------

// Add returns a + b
func (a Polar) Add(b Polar) Polar {
	return (a.PolarToCartesian().Add(b.PolarToCartesian())).CartesianToPolar()
}

//-----------------------------------------------------------------------------

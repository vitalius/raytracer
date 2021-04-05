package vector

import "math"

type Vector struct {
	X, Y, Z float64
}

func (r Vector) Add(s Vector) Vector {
	r.X += s.X
	r.Y += s.Y
	r.Z += s.Z
	return r
}

func (r Vector) Scale(t float64) Vector {
	r.X *= t
	r.Y *= t
	r.Z *= t
	return r
}

func (r Vector) Length() float64 {
	return math.Sqrt(r.X*r.X + r.Y*r.Y + r.Z*r.Z)
}

func (r Vector) Subtract(s Vector) Vector {
	r.X -= s.X
	r.Y -= s.Y
	r.Z -= s.Z
	return r
}

func (r Vector) Unit () Vector {
	len := r.Length()
	r.X /= len 
	r.Y /= len 
	r.Z /= len 
	return r
}

func (r Vector) Dot( a, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}
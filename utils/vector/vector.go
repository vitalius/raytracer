package vector

import "math"

type Vec3 struct {
    X, Y, Z float64
}

func (r Vec3) Add(s Vec3) Vec3 {
    r.X += s.X
    r.Y += s.Y
    r.Z += s.Z
    return r
}

func (r Vec3) Scale(t float64) Vec3 {
    r.X *= t
    r.Y *= t
    r.Z *= t
    return r
}

func (r Vec3) Length() float64 {
    return math.Sqrt(r.X*r.X + r.Y*r.Y + r.Z*r.Z)
}

func (r Vec3) Subtract(s Vec3) Vec3 {
    r.X -= s.X
    r.Y -= s.Y
    r.Z -= s.Z
    return r
}

func (r Vec3) Unit () Vec3 {
    len := r.Length()
    r.X /= len 
    r.Y /= len 
    r.Z /= len 
    return r
}

func (r Vec3) Dot( a, b Vec3) float64 {
    return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}
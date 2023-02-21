package vector

import "math"

type Vec3 struct {
    X, Y, Z float64
}

func BuildVec3 (x,y,z float64) Vec3 {
    var v Vec3
    v.X = x
    v.Y = y
    v.Z = z
    return v
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

func (r Vec3) Dot(a, b Vec3) float64 {
    return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (r Vec3) ToRgb() [3]byte {
    result := [3]byte {0x0, 0x0, 0x0}
    result[0] = byte( 255.99*r.X )
    result[1] = byte( 255.99*r.Y )
    result[2] = byte( 255.99*r.Z )
    return result
}

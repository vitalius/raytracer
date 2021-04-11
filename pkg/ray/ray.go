package ray

import (
	"math"
	. "raytracer/pkg/object"
	. "raytracer/utils/vector"
)

type Ray struct {
	A, B Vec3
}

func (r Ray) Origin() Vec3 {
	return r.A
}

func (r Ray) Direction() Vec3 {
	return r.B
}

func (r Ray) Point(t float64) Vec3 {
	s := r.B.Scale(t)
	return r.A.Add(s)
}

func (r *Ray) Color(sphere Sphere) Vec3 {
    hit := r.Hit(sphere)
    if hit > 0.0 {
        pat := r.Point(hit)
        pat = pat.Subtract(BuildVec3(0,0,-1))
        n := pat.Unit()        
        res := BuildVec3(n.X+1.0, n.Y+1.0, n.Z+1.0)
        return res.Scale(0.5)
    }

    direction := r.Direction()
    direction = direction.Unit()

    t := 0.5*(direction.Y + 1.0)
    
    start := BuildVec3(0.5, 0.7, 1.0)
    start = start.Scale(1.0-t)

    end := BuildVec3(1.0, 1.0, 1.0)
    end = end.Scale(t)

    result:=start.Add(end)
    return result
}

func (r *Ray) Hit(sphere Sphere) float64 {
	oc := r.Origin()
	oc = oc.Subtract(sphere.Center)
	a := oc.Dot(r.Direction(), r.Direction())
	b := 2.0 * oc.Dot(oc, r.Direction())
	c := oc.Dot(oc, oc) - sphere.Radius*sphere.Radius
	disc := b*b - 4*a*c
	if disc < 0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(disc)) / (2.0 * a)
	}
}
package ray

import (
	"math"
	. "raytracer/utils/vector"
)

type Ray struct {
	A, B Vec3
}

func (r Ray) origin() Vec3 {
	return r.A
}

func (r Ray) dir() Vec3 {
	return r.B
}

func (r Ray) Point(t float64) Vec3 {
	s := r.B.Scale(t)
	return r.A.Add(s)
}


func (r *Ray) Color() Vec3 {
    sphere_c := Vec3{0.0, 0.0, -1.0}
    sphere_r := float64(0.5)
    hit := r.hit_sphere(sphere_c, sphere_r)
    if hit > 0.0 {
        pat := r.Point(hit)
        pat = pat.Subtract(Vec3{0,0,-1})
        n := pat.Unit()        
        res := Vec3{n.X+1.0, n.Y+1.0, n.Z+1.0}
        return res.Scale(0.5)
    }

    direction := r.dir()
    direction = direction.Unit()

    t := 0.5*(direction.Y + 1.0)
    
    start := Vec3{0.5, 0.7, 1.0}
    start = start.Scale(1.0-t)

    end := Vec3{1.0, 1.0, 1.0}
    end = end.Scale(t)

    result:=start.Add(end)
    return result
}

func (r Ray) hit_sphere(center Vec3, radius float64) float64 {
    oc := r.origin()
    oc = oc.Subtract(center)
    a := oc.Dot(r.dir(), r.dir())
    b := 2.0 * oc.Dot(oc, r.dir())
    c := oc.Dot(oc, oc) - radius*radius
    disc := b*b - 4*a*c
    if disc < 0 {
        return -1.0
    } else {
        return (-b - math.Sqrt(disc)) / (2.0*a)
    }
}
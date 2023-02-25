package ray

import (
    //"fmt"
	"math"
	. "raytracer/pkg/object"
	. "raytracer/utils/vector"
)

type HitRecord struct {
    p, normal Vec3
    t float64
}

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

func (r *Ray) Color(spheres []Sphere) Vec3 {
    hr := HitRecord {p:BuildVec3(0,0,0), normal:BuildVec3(0,0,0), t:0.0}
    hitSomething := false
    closest := hr
    closest.t = 10000

    for _, s := range spheres {     
        hit := r.Hit(s, &hr)
        if hit && hr.t < closest.t {
            closest = hr
            hitSomething = true
        }
    }
    if (hitSomething) {
        // generate some color values
        return BuildVec3(1,1,1).Add(closest.normal).Scale(0.5)
    }

    /* No hit, draw sky */
    direction := r.Direction().Unit()
    t := 0.5*(direction.Y + 1.0)
    start := BuildVec3(0.5, 0.7, 1.0).Scale(1.0-t)
    end := BuildVec3(1.0, 1.0, 1.0).Scale(t)
    return start.Add(end)
}

func (r *Ray) Hit(sphere Sphere, hr *HitRecord) bool {
	oc := r.Origin().Subtract(sphere.Center)
	
    a := r.Direction().Length();
    a = a*a

	b := oc.Dot(oc, r.Direction())
	
    c := oc.Length()
    c = c*c
    c = c - sphere.Radius*sphere.Radius
	
    disc := b*b - a*c

    // no hit
	if disc < 0 {
		return false
	}

    // find roots
    sqrt_disc := math.Sqrt(disc)
	root1 := (-b - sqrt_disc) / a
    root2 := (-b + sqrt_disc) / a
    root := math.Min(root1, root2)

    // check if root is nearest or valid
    if (root < hr.t) {
        return false;
    }

    // set hit record values
    hr.t = root
    hr.p = r.Point(hr.t)
    hr.normal = (hr.p.Subtract(sphere.Center)).Scale(1.0/sphere.Radius)

    // check for normal direction, make sure its outward
    if (oc.Dot(r.Direction(), hr.normal)) > 0 {
        hr.normal = hr.normal.Scale(-1)
    }

    return true
}
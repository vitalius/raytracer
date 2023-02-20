package object

import (
	. "raytracer/utils/vector"
)

type SceneObject interface {
	Hit() float64
}

type Sphere struct {
    Center Vec3 `json:"center"`
    Radius float64 `json:"radius"`
}

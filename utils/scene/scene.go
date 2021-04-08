package scene

import (
	"encoding/json"
	"log"

	. "raytracer/utils/vector"
)

type Sphere struct {
    Center Vec3 `json:"center"`
    Radius float64 `json:"radius"`
}

type Scene struct {
    Raster struct {
        Width  int `json:"width"`
        Height int `json:"height"`
    } `json:"raster"`

    Camera struct {
        Location  Vec3 `json:"location"`
        Direction Vec3 `json:"direction"`
    } `json:"camera"`

    Screen struct {
        LowerLeft  Vec3 `json:"lower_left"`
        UpperRight Vec3 `json:"upper_right"`
    } `json:"screen"`

    Spheres []Sphere `json:"spheres"`
}

func LoadFromJson(jsonBytes []byte) Scene {
    var scene Scene
	marshal_err := json.Unmarshal(jsonBytes, &scene)
	if marshal_err != nil {
		log.Fatal(marshal_err)
	}
	return scene
}
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"

	. "raytracer/utils/input"
	. "raytracer/utils/raster"
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

func (r Ray) point(t float64) Vec3 {
    s := r.B.Scale(t)
    return r.A.Add(s)
}

func LoadSceneFile (filename string) Scene {
    jsonFile, open_err := os.Open(filename)
    if open_err != nil {
		log.Fatal(open_err)
    }

    defer jsonFile.Close()

    jsonBytes, read_err := ioutil.ReadAll(jsonFile)
    if read_err != nil {
		log.Fatal(read_err)
    }

    var scene = LoadFromJson(jsonBytes)
    return scene
}

func (r Ray) color() Vec3 {
    sphere_c := Vec3{0.0, 0.0, -1.0}
    sphere_r := float64(0.5)
    hit := r.hit_sphere(sphere_c, sphere_r)
    if hit > 0.0 {
        pat := r.point(hit)
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

func main() {

    // no scene file, show help message
    if len(os.Args) < 2 {
        fmt.Printf("Specify scene description file (JSON)\n")
        fmt.Printf("\t%s <scene_file.json>\n", filepath.Base(os.Args[0]) )
        return;
    }

    // load scene file
    var scene = LoadSceneFile(os.Args[1])

    img_width := scene.Raster.Width
    img_height:= scene.Raster.Height
    b := NewBitmap(img_width, img_height)

    llc_o := scene.Screen.LowerLeft
    hor_o := Vec3{4.0, 0.0, 0.0}
    ver_o := Vec3{0.0, 2.0, 0.0}
    org_o := Vec3{0.0, 0.0, 0.0}

    for y := img_height - 1; y >= 0; y-- {
        for x := 0; x < img_width; x++ {
            u := float64(x) / float64(img_width)
            v := float64(y) / float64(img_height)
            hor := hor_o.Scale(u)
            ver := ver_o.Scale(v)
            hor = hor.Add(ver)
            llc := llc_o.Add(hor)

            ray := Ray{org_o, llc}
            c := ray.color()

            red := byte( 255.99*c.X )
            green := byte( 255.99*c.Y )
            blue := byte(255.99*c.Z)

            b.SetPx(x,y,Pixel{R:red, G:green, B:blue});
        }
    }

    err := b.WritePngFile("output.png")
    if err != nil {
        fmt.Println(err)
    }
}

 
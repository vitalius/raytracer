package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	. "raytracer/utils/raster"
	. "raytracer/utils/ray"
	. "raytracer/utils/scene"
	. "raytracer/utils/vector"
)


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
    hor_o := BuildVec3(4.0, 0.0, 0.0)
    ver_o := BuildVec3(0.0, 2.0, 0.0)
    org_o := BuildVec3(0.0, 0.0, 0.0)

    for y := img_height - 1; y >= 0; y-- {
        for x := 0; x < img_width; x++ {
            u := float64(x) / float64(img_width)
            v := float64(y) / float64(img_height)
            hor := hor_o.Scale(u)
            ver := ver_o.Scale(v)
            hor = hor.Add(ver)
            llc := llc_o.Add(hor)

            ray := Ray {A:org_o, B:llc}
            c := ray.Color()

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

 
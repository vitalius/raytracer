package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"sync"
	"time"

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


func YWorkSplit(start, end, count int) [][]int {
    var bucket_size = int(math.Ceil(float64(end-start)/float64(count)))
    var result = make([][]int, 0)
    if (bucket_size < 1) {
        var use_all = []int{start,end}
        return append(result, use_all)
    }
    var i = end;
    
    for i > bucket_size {
        i -= bucket_size
        result = append(result, []int{i+bucket_size, i})
    }
    result = append(result, []int{i,start})

    return result
}


func worker(y []int, wg *sync.WaitGroup) {
    defer wg.Done()    
    fmt.Printf("Worker %d starting\n", y)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", y)
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

    var y_split = YWorkSplit(0, img_height, 8)

    var wg sync.WaitGroup

    for _, bucket := range y_split {
        wg.Add(1)
        go func(y1, y2, img_width int, b *Bitmap, hor_o, ver_o, llc_o, org_o Vec3) {
            defer wg.Done()
            for y := y1 - 1; y >= y2; y-- {
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

                        b.SetPx(x,y,Pixel{R:red, G:green, B:blue})
                    }
                }
            }(bucket[0], bucket[1], img_width, b, hor_o, ver_o, llc_o, org_o)
        }
    wg.Wait()

    err := b.WritePngFile("output.png")
    if err != nil {
        fmt.Println(err)
    }
}

 
package raster

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

type Rgb int32

type Pixel struct {
    R, G, B byte
}

func (c Rgb) Pixel() Pixel {
    return Pixel{R: byte(c >> 16), G: byte(c >> 8), B: byte(c)}
}

func (p Pixel) Rgb() Rgb {
    return Rgb(p.R)<<16 | Rgb(p.G)<<8 | Rgb(p.B)
}

type Bitmap struct {
    Comments   []string
    rows, cols int
    px         []Pixel
    pxRow      [][]Pixel
}

func NewBitmap(x, y int) (b *Bitmap) {
    b = &Bitmap{
        rows:     y, 
        cols:     x,
        px:       make([]Pixel, x*y),
        pxRow:    make([][]Pixel, y),
    }
    x0, x1 := 0, x
    for i := range b.pxRow {
        b.pxRow[i] = b.px[x0:x1] 
        x0, x1 = x1, x1+x
    }
    return b
}

func (b *Bitmap) SetPx(x, y int, p Pixel) bool {
    defer func() { recover() }()
    b.pxRow[y][x] = p
    return true
}

func (b *Bitmap) GetPx(x, y int) (p Pixel, ok bool) {
    defer func() { recover() }()
    return b.pxRow[y][x], true
}

func (b *Bitmap) GetPxRgb(x, y int) (Rgb, bool) {
    p, ok := b.GetPx(x, y)
    if !ok {
        return 0, false
    }
    return p.Rgb(), true
}

func (b *Bitmap) WritePngTo(w io.Writer) (img_result *image.RGBA, err error) {
    width := b.cols;
    height:= b.rows;
    upLeft := image.Point{0, 0}
    lowRight := image.Point{width, height}
    img_result = image.NewRGBA(image.Rectangle{upLeft, lowRight})

    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            rgb, _ := b.GetPxRgb(x,y)
            px := rgb.Pixel()
            img_result.Set(x, y, color.RGBA{px.R, px.G, px.B, 0xff})
        }
    }
    return img_result, nil
}

func (b *Bitmap) WritePngFile(fn string) (err error) {
    var f *os.File
    f, err = os.Create(fn);
    if err != nil {
        return
    }
    img, err := b.WritePngTo(f);
    if err != nil {
        return
    }
    png.Encode(f, img)
    return f.Close()
}

package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func resizeBilinear(img image.Image, width, height int) image.Image {
	tmp := image.NewRGBA(image.Rect(0, 0, width, height))
	oldWidth, oldHeight := img.Bounds().Dx(), img.Bounds().Dy()
	x_ratio := float64(oldWidth) / float64(width)
	y_ratio := float64(oldHeight) / float64(height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			x := int(x_ratio * float64(j))
			y := int(y_ratio * float64(i))
			x_diff := x_ratio*float64(j) - float64(x)
			y_diff := y_ratio*float64(i) - float64(y)
			r, g, b, a := img.At(x, y).RGBA()
			r1, g1, b1, a1 := img.At(x, y+1).RGBA()
			r2, g2, b2, a2 := img.At(x+1, y).RGBA()
			r3, g3, b3, a3 := img.At(x+1, y+1).RGBA()

			// red element
			// Yr = Ar(1-w)(1-h) + Br(w)(1-h) + Cr(h)(1-w) + Dr(wh)
			red := float64(r>>8)*(1-x_diff)*(1-y_diff) + float64(r1>>8)*(
				x_diff)*(1-y_diff) + float64(r2>>8)*(y_diff)*(1-x_diff) + float64(r3>>8)*(x_diff*y_diff)
			// green element
			// Yg = Ag(1-w)(1-h) + Bg(w)(1-h) + Cg(h)(1-w) + Dg(wh)
			green := float64(g>>8)*(1-x_diff)*(1-y_diff) + float64(g1>>8)*(x_diff)*(1-y_diff) + float64(g2>>8)*(y_diff)*(1-x_diff) + float64(g3>>8)*(x_diff*y_diff)
			// blue element
			// Yr = Ar(1-w)(1-h) + Br(w)(1-h) + Cr(h)(1-w) + Dr(wh)
			blue := float64(b>>8)*(1-x_diff)*(1-y_diff) + float64(b1>>8)*(x_diff)*(1-y_diff) + float64(b2>>8)*(y_diff)*(1-x_diff) + float64(b3>>8)*(x_diff*y_diff)

			alpha := float64(a>>8)*(1-x_diff)*(1-y_diff) + float64(a1>>8)*(x_diff)*(1-y_diff) + float64(a2>>8)*(y_diff)*(1-x_diff) + float64(a3>>8)*(x_diff*y_diff)

			col := color.RGBA{
				R: uint8(red),
				G: uint8(green),
				B: uint8(blue),
				A: uint8(alpha),
			}
			tmp.Set(j, i, col)
		}
	}
	return tmp
}

func main() {
	in, err := os.Open("./grape.png")
	if err != nil {
		log.Fatal("open file err:", err)
	}
	defer in.Close()

	img, _, err := image.Decode(in)
	if err != nil {
		log.Fatal("decode file err:", err)
	}
	out, err := os.Create("./grape1.png")
	if err != nil {
		log.Fatal("create file err:", err)
	}
	defer out.Close()

	result := resizeBilinear(img, 162*4, 179*4)
	_ = png.Encode(out, result)
}

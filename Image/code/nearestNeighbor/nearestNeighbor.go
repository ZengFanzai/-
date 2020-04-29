package main

import (
	"image"
	"image/png"
	_ "image/png"
	"log"
	"math"
	"os"
)

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

	result := resize1(img, 162*4, 179*4)
	_ = png.Encode(out, result)
}

func resize(img image.Image, width, height int) image.Image {
	tmp := image.NewRGBA(image.Rect(0, 0, width, height))
	x_ratio := float64(img.Bounds().Dx()) / float64(width)
	y_ratio := float64(img.Bounds().Dy()) / float64(height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			px := math.Floor(float64(j) * x_ratio)
			py := math.Floor(float64(i) * y_ratio)
			tmp.Set(j, i, img.At(int(px), int(py)))
		}
	}
	return tmp
}

func resize1(img image.Image, width, height int) image.Image {
	tmp := image.NewRGBA(image.Rect(0, 0, width, height))
	x_ratio := img.Bounds().Dx()<<16/width + 1
	y_ratio := img.Bounds().Dy()<<16/height + 1
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			px := (j * x_ratio) >> 16
			py := (i * y_ratio) >> 16
			tmp.Set(j, i, img.At(px, py))
		}
	}
	return tmp
}

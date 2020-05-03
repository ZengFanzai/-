package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func trilinearImageScaling(larger, smaller image.Image, width,
	height int) image.Image {
	tmp := image.NewRGBA(image.Rect(0, 0, width, height))
	lw, lh := larger.Bounds().Dx(), larger.Bounds().Dy()
	sw, sh := smaller.Bounds().Dx(), smaller.Bounds().Dy()
	lw_ratio := float64(lw-1) / float64(width)
	lh_ratio := float64(lh-1) / float64(height)
	sw_ratio := float64(sw-1) / float64(width)
	sh_ratio := float64(sh-1) / float64(height)
	// 估算h3的距离
	h3_diff := float64(lw-width) / float64(lw-sw)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// larger
			lx := int(lw_ratio * float64(j))
			ly := int(lh_ratio * float64(i))
			lw_diff := lw_ratio*float64(j) - float64(lx)
			lh_diff := lh_ratio*float64(i) - float64(ly)
			lr, lg, lb, la := larger.At(lx, ly).RGBA()
			lr1, lg1, lb1, la1 := larger.At(lx, ly+1).RGBA()
			lr2, lg2, lb2, la2 := larger.At(lx+1, ly).RGBA()
			lr3, lg3, lb3, la3 := larger.At(lx+1, ly+1).RGBA()

			// smaller
			sx := int(sw_ratio * float64(j))
			sy := int(sh_ratio * float64(i))
			sw_diff := sw_ratio*float64(j) - float64(sx)
			sh_diff := sh_ratio*float64(i) - float64(sy)
			sr, sg, sb, sa := larger.At(sx, sy).RGBA()
			sr1, sg1, sb1, sa1 := larger.At(sx, sy+1).RGBA()
			sr2, sg2, sb2, sa2 := larger.At(sx+1, sy).RGBA()
			sr3, sg3, sb3, sa3 := larger.At(sx+1, sy+1).RGBA()

			blue := float64(lb>>8)*(1-lw_diff)*(1-lh_diff)*(1-h3_diff) +
				float64(lb1>>8)*lw_diff*(1-lh_diff)*(1-h3_diff) +
				float64(lb2>>8)*(1-lw_diff)*lh_diff*(1-h3_diff) +
				float64(lb3>>8)*lw_diff*lh_diff*(1-h3_diff) +
				float64(sb>>8)*(1-sw_diff)*(1-sh_diff)*h3_diff +
				float64(sb1>>8)*sw_diff*(1-sh_diff)*h3_diff +
				float64(sb2>>8)*(1-lw_diff)*sh_diff*h3_diff +
				float64(sb3>>8)*lw_diff*sh_diff*h3_diff

			green := float64(lg>>8)*(1-lw_diff)*(1-lh_diff)*(1-h3_diff) +
				float64(lg1>>8)*lw_diff*(1-lh_diff)*(1-h3_diff) +
				float64(lg2>>8)*(1-lw_diff)*lh_diff*(1-h3_diff) +
				float64(lg3>>8)*lw_diff*lh_diff*(1-h3_diff) +
				float64(sg>>8)*(1-sw_diff)*(1-sh_diff)*h3_diff +
				float64(sg1>>8)*sw_diff*(1-sh_diff)*h3_diff +
				float64(sg2>>8)*(1-lw_diff)*sh_diff*h3_diff +
				float64(sg3>>8)*lw_diff*sh_diff*h3_diff

			red := float64(lr>>8)*(1-lw_diff)*(1-lh_diff)*(1-h3_diff) +
				float64(lr1>>8)*lw_diff*(1-lh_diff)*(1-h3_diff) +
				float64(lr2>>8)*(1-lw_diff)*lh_diff*(1-h3_diff) +
				float64(lr3>>8)*lw_diff*lh_diff*(1-h3_diff) +
				float64(sr>>8)*(1-sw_diff)*(1-sh_diff)*h3_diff +
				float64(sr1>>8)*sw_diff*(1-sh_diff)*h3_diff +
				float64(sr2>>8)*(1-lw_diff)*sh_diff*h3_diff +
				float64(sr3>>8)*lw_diff*sh_diff*h3_diff

			alpha := float64(la>>8)*(1-lw_diff)*(1-lh_diff)*(1-h3_diff) +
				float64(la1>>8)*lw_diff*(1-lh_diff)*(1-h3_diff) +
				float64(la2>>8)*(1-lw_diff)*lh_diff*(1-h3_diff) +
				float64(la3>>8)*lw_diff*lh_diff*(1-h3_diff) +
				float64(sa>>8)*(1-sw_diff)*(1-sh_diff)*h3_diff +
				float64(sa1>>8)*sw_diff*(1-sh_diff)*h3_diff +
				float64(sa2>>8)*(1-lw_diff)*sh_diff*h3_diff +
				float64(sa3>>8)*lw_diff*sh_diff*h3_diff

			col := color.RGBA{
				R: uint8(red),
				G: uint8(green),
				B: uint8(blue),
				A: uint8(alpha),
			}
			tmp.Set(j,i, col)
		}
	}
	return tmp
}

func main() {
	in1, err := os.Open("./larger.png")
	if err != nil {
		log.Fatal("open file err:", err)
	}
	defer in1.Close()
	img1,err := png.Decode(in1)
	if err != nil {
		log.Fatal(err)
	}

	in2, err := os.Open("./smaller.png")
	if err != nil {
		log.Fatal("open file err:", err)
	}
	defer in2.Close()
	img2,err := png.Decode(in2)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("./mid.png")
	if err != nil {
		log.Fatal("create file err:", err)
	}
	defer out.Close()
	result := trilinearImageScaling(img1, img2, 200, 200)
	_ = png.Encode(out, result)
}

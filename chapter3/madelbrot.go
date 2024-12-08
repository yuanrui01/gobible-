package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 2048, 2048
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 点(px,py) 表示复数值z
			img.Set(px, py, madelbrot(z))
		}
	}
	// 创建一个文件来保存图像
	file, err := os.Create("chapter3/mandelbrot2.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 将图像编码为PNG并保存到文件
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func madelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

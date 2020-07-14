package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{
	width int
	height int
	pixels [][]uint8
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := img.pixels[y][x]
	return color.RGBA{v, v, 123, 200}
}

func Pic(dx,dy int)[][]uint8{
	img := make([][]uint8,dy)
	for i:=0;i<dy;i+=1{
		img[i] = make([]uint8,dx)
	}
	for i:=0;i<dy;i+=1{
		for j:=0;j<dx;j+=1{
			img[i][j]=(uint8)(i^j)/2
		}
	}
	return img
}

func main() {
	m := Image{256,256,Pic(256,256)}
	pic.ShowImage(m)
}

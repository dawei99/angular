//
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main(){

	size := 300
	pic := image.NewGray(image.Rect(0,0,size, size))
	for xi:=0; xi<size; xi++ {
		for yi:=0; yi<size; yi++ {
			pic.SetGray(xi, yi, color.Gray{255})
		}
	}
	file, err := os.Create("sin.png")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	height:=50
	yStart:=100
	for gyi:=0; gyi<height; gyi++ {
		yPoint := yStart + gyi
		for i:=50; i<250; i++ {
			pic.SetGray(i, yPoint, color.Gray{200})
		}
	}
	png.Encode(file, pic)

}
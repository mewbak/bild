package main

import (
	"image/png"
	"os"

	"github.com/roz3x/bild/noise"
)

func main() {
	img := noise.PerlinGenerate(200, 200, 0.2)
	f, _ := os.Create("file.png")
	png.Encode(f, img)
}

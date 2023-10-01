package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // import this package to decode JPEGs
	_ "image/png"  // import this package to decode PNGs
	"os"
)

// PrintImagePixels prints the RGB values of the pixels in the provided image.
func PrintImagePixels(filename string) error {
	reader, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			fmt.Printf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
		}
	}

	return nil
}

package img_colors

import (
	"fmt"
	"image"
	_ "image/jpeg" // Import this package to decode JPEGs
	_ "image/png"  // Import this package to decode PNGs
	"os"
)

// PrintImagePixelsToFile prints the RGB values of the pixels in the provided image to a file.
// It takes a filename for the output file and a filename for the input image as input and returns an error if there are any issues.
func PrintImagePixelsToFile(inputFilename, outputFilename string) error {
	reader, err := os.Open(inputFilename)
	if err != nil {
		return fmt.Errorf("failed to open input file: %v", err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			// Write the output to the file instead of printing to the console.
			fmt.Fprintf(outputFile, "Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
		}
	}

	return nil
}

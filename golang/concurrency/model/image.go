package model

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
	"time"

	"github.com/nfnt/resize"
)

type Image struct {
	Name       string
	Image      image.Image
	InputPath  string
	OutputPath string
}

type Images []Image

func PrepareImages(paths []string, outPath string) <-chan Image {
	out := make(chan Image)
	go func() {
		defer close(out)

		for _, p := range paths {
			subStr := strings.Split(p, "/")

			img := Image{
				Name:       subStr[len(subStr)-1],
				InputPath:  p,
				OutputPath: strings.Replace(p, "images/", outPath, 1),
			}

			img.Image = readImage(p)
			out <- img
		}
	}()
	return out
}

func readImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println(path)
		panic(err)
	}

	return img
}

func Write(path string, img image.Image) {
	outputFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Encode the image to the new file
	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 1)
}

func Grayscale(path string, img image.Image) image.Image {
	// Create a new grayscale image
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}

	return grayImg
}

func Resize(path string, img image.Image) image.Image {
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	time.Sleep(time.Second * 5)

	return resizedImg
}

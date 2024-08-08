package pipeline

import (
	"fmt"

	"github.com/micamics/atelier/golang/concurrency/model"
)

func ExecutePipeline() {
	imagePaths := []string{
		"images/image1.png",
		"images/image2.png",
		"images/image3.png",
		"images/image4.png",
	}

	imgChan := model.PrepareImages(imagePaths, "pipeline/output/")

	resizedImgChan := resizeImages(imgChan)
	grayscaledImgChan := grayscaleImages(resizedImgChan)
	writtenImgChan := writeImages(grayscaledImgChan)

	for img := range writtenImgChan {
		fmt.Println("processed image: ", img.Name)
	}
}

func resizeImages(in <-chan model.Image) <-chan model.Image {
	out := make(chan model.Image)

	go func() {
		defer close(out)
		for img := range in {
			img.Image = model.Resize(img.InputPath, img.Image)
			out <- img
		}
	}()

	return out
}

func grayscaleImages(in <-chan model.Image) <-chan model.Image {
	out := make(chan model.Image)

	go func() {
		defer close(out)
		for img := range in {
			img.Image = model.Grayscale(img.InputPath, img.Image)
			out <- img
		}
	}()

	return out
}

func writeImages(in <-chan model.Image) <-chan model.Image {
	out := make(chan model.Image)
	go func() {
		defer close(out)

		for img := range in {
			model.Write(img.OutputPath, img.Image)
			out <- img
		}
	}()

	return out
}

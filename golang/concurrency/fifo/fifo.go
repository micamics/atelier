package fifo

import (
	"fmt"
	"sync"

	"github.com/micamics/atelier/golang/concurrency/model"
)

func ExecuteFanInFanOut() {
	numWorkers := 4
	doneChan := make(chan bool)
	defer close(doneChan)

	imagePaths := []string{
		"images/image1.png",
		"images/image2.png",
		"images/image3.png",
		"images/image4.png",
	}

	imgChan := model.PrepareImages(imagePaths, "fifo/output/")
	workers := make([]<-chan model.Image, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workers[i] = resizeImages(doneChan, imgChan)
	}

	for i := 0; i < numWorkers; i++ {
		workers[i] = grayscaleImages(doneChan, fanIn(doneChan, workers...))
	}

	processedImages := writeImages(doneChan, fanIn(doneChan, workers...))
	for img := range processedImages {
		fmt.Println("processed image: ", img.Name)
	}
}

func resizeImages(doneChan <-chan bool, in <-chan model.Image) <-chan model.Image {
	out := make(chan model.Image)

	go func() {
		defer close(out)
		for img := range orDone(doneChan, in) {
			img.Image = model.Resize(img.InputPath, img.Image)
			out <- img
		}
	}()

	return out
}

func grayscaleImages(doneChan <-chan bool, in <-chan model.Image) <-chan model.Image {
	out := make(chan model.Image)

	go func() {
		defer close(out)
		for img := range orDone(doneChan, in) {
			img.Image = model.Grayscale(img.InputPath, img.Image)
			out <- img
		}
	}()

	return out
}

func writeImages(doneChan <-chan bool, in <-chan model.Image) <-chan model.Image {
	out := make(chan model.Image)

	go func() {
		defer close(out)
		for img := range orDone(doneChan, in) {
			model.Write(img.OutputPath, img.Image)
			out <- img
		}
	}()

	return out
}

func fanIn(doneChan <-chan bool, fannedOutChannels ...<-chan model.Image) <-chan model.Image {
	var wg sync.WaitGroup

	wg.Add(len(fannedOutChannels))
	outImages := make(chan model.Image)
	multiplex := func(ch <-chan model.Image) {
		defer wg.Done()
		for i := range ch {
			select {
			case <-doneChan:
				return
			case outImages <- i:
			}
		}
	}

	for _, ch := range fannedOutChannels {
		go multiplex(ch)
	}
	go func() {
		wg.Wait()
		close(outImages)
	}()

	return outImages
}

func orDone(done <-chan bool, img <-chan model.Image) <-chan model.Image {
	relayStream := make(chan model.Image)

	go func() {
		defer close(relayStream)

		for {
			select {
			case <-done:
				return
			case v, ok := <-img:
				if !ok {
					return
				}

				select {
				case relayStream <- v:
				case <-done:
					return

				}
			}
		}
	}()

	return relayStream
}

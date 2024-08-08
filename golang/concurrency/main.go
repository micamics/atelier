package main

import (
	"fmt"
	"time"

	"github.com/micamics/atelier/golang/concurrency/fifo"
	"github.com/micamics/atelier/golang/concurrency/pipeline"
)

func main() {
	var Yellow = "\033[33m"
	var Reset = "\033[0m"

	fmt.Printf("============= " + Yellow + "PIPELINE" + Reset + " =============\n")
	start := time.Now()
	pipeline.ExecutePipeline()
	took := time.Since(start).Seconds()

	fmt.Printf("pipeline took "+Yellow+"%f"+Reset+" seconds to complete\n", took)
	fmt.Println("====================================")

	fmt.Printf("\n============== " + Yellow + "FIFO" + Reset + " ================\n")
	startFifo := time.Now()
	fifo.ExecuteFanInFanOut()
	took = time.Since(startFifo).Seconds()

	fmt.Printf("fifo took "+Yellow+"%f"+Reset+" seconds to complete\n", took)
	fmt.Println("====================================\n\n")
}

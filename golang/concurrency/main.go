package main

import (
	"fmt"
	"time"

	"github.com/micamics/atelier/golang/concurrency/fofi"
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

	fmt.Printf("\n========= " + Yellow + "FAN-OUT,FAN-IN" + Reset + " ===========\n")

	startFifo := time.Now()
	fofi.ExecuteFanOutFanIn()
	took = time.Since(startFifo).Seconds()

	fmt.Printf("fan-out, fan-in took "+Yellow+"%f"+Reset+" seconds to complete\n", took)
}

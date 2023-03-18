package main

import (
	"Formatter/phase2"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const (
	defaultNWorker    = 5
	defaultInputFile  = "input.txt"
	defaultOutputFile = "output.txt"
)

var nWorker = flag.Int("n", defaultNWorker, "number of workers")
var inputFile = flag.String("input", defaultInputFile, "input file")
var outputFile = flag.String("output", defaultOutputFile, "output file")

func createWorkerPool(nWorker int, tasks <-chan *phase2.Task, taskResults chan<- *phase2.TaskResult) {
	var wg sync.WaitGroup
	wg.Add(nWorker)
	for i := 0; i < nWorker; i++ {
		go func() {
			phase2.Worker(tasks, taskResults, &wg)
		}()
	}
	wg.Wait()
}

func main() {
	flag.Parse()

	input, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(input)

	output, err := os.Create(*outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func(output *os.File) {
		_ = output.Close()
	}(output)

	startTime := time.Now()
	scanner := bufio.NewScanner(input)
	coordinator, taskCount, err := phase2.MakeCoordinator(scanner)
	if err != nil {
		panic(err)
	}

	tasks := make(chan *phase2.Task, taskCount)
	taskResults := make(chan *phase2.TaskResult, taskCount)

	// Start worker pool, allocate tasks, and wait for tasks to be completed
	go createWorkerPool(*nWorker, tasks, taskResults)
	go coordinator.Allocate(tasks)

	err = coordinator.HandleResult(output, taskResults)
	close(taskResults)
	if err != nil {
		panic(err)
	}

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken:", diff.Seconds(), "seconds")
}

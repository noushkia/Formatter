package main

import (
	"Formatter/phase2"
	"flag"
	"fmt"
	"sync"
	"time"
)

func createWorkerPool(nWorker int, tasks chan *phase2.Task, taskResults chan *phase2.TaskResult) {
	var wg sync.WaitGroup

	wg.Add(nWorker)
	for i := 0; i < nWorker; i++ {
		go func() {
			err := phase2.Worker(&wg, tasks, taskResults)
			if err != nil {
				println("Error in worker")
			}
		}()
	}
	wg.Wait()
}

func main() {
	var nWorker int
	var inputFile string
	flag.IntVar(&nWorker, "n", 5, "number of workers")
	flag.StringVar(&inputFile, "file", "input.txt", "input file")
	flag.Parse()
	coordinator, tasksCount := phase2.MakeCoordinator(inputFile)

	// initialize the channels
	var tasks = make(chan *phase2.Task, tasksCount)             // send to channel
	var taskResults = make(chan *phase2.TaskResult, tasksCount) // receive from this channel

	startTime := time.Now()

	go func() {
		err := coordinator.Allocate(tasks)
		if err != nil {
			println("Error in allocation")
		}
		close(tasks)
	}()

	go createWorkerPool(nWorker, tasks, taskResults)

	err := coordinator.HandleResult(taskResults)
	if err != nil {
		panic(err)
	}

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken:", diff.Seconds(), "seconds")
}

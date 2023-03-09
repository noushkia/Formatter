package main

import (
	"Formatter/phase2"
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
	close(taskResults)
}

func main() {
	startTime := time.Now()
	nWorker := 5
	coordinator := phase2.MakeCoordinator("input.txt", nWorker) // get from cla
	// initialize the channels
	var tasks = make(chan *phase2.Task, nWorker)             // send to channel
	var taskResults = make(chan *phase2.TaskResult, nWorker) // receive from this channel

	go func() {
		err := coordinator.Allocate(tasks)
		if err != nil {
			println("Error in allocation")
		}
	}()
	createWorkerPool(nWorker, tasks, taskResults) // get from cla
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

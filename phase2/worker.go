package phase2

import (
	"fmt"
	"sync"
	"time"
)

func Worker(wg *sync.WaitGroup, tasks chan *Task, taskResults chan *TaskResult) error {
	// use select to get the jobs from each channel?
	select {
	case task := <-tasks:
		output := &TaskResult{id: task.id, line: task.line} // todo: format the line
		taskResults <- output
	case taskResult := <-taskResults:
		fmt.Printf(string(rune(taskResult.id)), taskResult.line) //todo: write to file (don't forget a mutex!)
	case <-time.After(time.Second * 1):
		fmt.Println("idle worker")
	}
	wg.Done()
	return nil
}

package main

import (
	"Formatter/phase1"
	"fmt"
	"time"
)

func main() {
	// todo: read from stdio
	inputSentence := "this is my 1 attempt to write a go program. hello world!"
	startTime := time.Now()
	println(inputSentence)
	println(phase1.Format(inputSentence))
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

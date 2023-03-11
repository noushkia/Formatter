package main

import (
	"Formatter/phase1"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(outputFile)

	startTime := time.Now()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputSentence := scanner.Text()

		formattedSentence := phase1.Format(inputSentence)

		_, err = fmt.Fprintln(outputFile, formattedSentence)
		if err != nil {
			log.Fatal(err)
		}
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken:", diff.Seconds(), "seconds")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

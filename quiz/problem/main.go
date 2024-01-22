package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// open CSV file
	csvFlag := flag.String("csv", "csv/problems.csv", "a csv file in the format of 'question,answer'")
	durationFlag := flag.Int("duration", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, error := os.Open(*csvFlag)
	if error != nil {
		fmt.Println(error)
	}
	if file != nil {
		defer file.Close()

		// read CSV file
		fileReader := csv.NewReader(file)
		records, error := fileReader.ReadAll()
		if error != nil {
			fmt.Println(error)
		}

		correctAnswers := 0
		var answer string

		go func() {
			time.Sleep(time.Duration(*durationFlag) * time.Second)
			fmt.Println("Time's up!")
			fmt.Println("You scored", correctAnswers, "out of", len(records))
			os.Exit(1)
		}()

		for _, record := range records {
			fmt.Println("What is", record[0], "?")
			fmt.Scanln(&answer)
			if answer == record[1] {
				correctAnswers++
			}
		}
		fmt.Println("You scored", correctAnswers, "out of", len(records))
	}
}

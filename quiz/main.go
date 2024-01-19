package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	// open CSV file
	args := os.Args
	args[1] = "30"
	duration, err := strconv.ParseInt(args[1], 10, 64)

	if err != nil {
		fmt.Println("Unable to parse duration, using default 30")
		duration = 30
	}
	file, error := os.Open("csv/problems.csv")
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
			time.Sleep(time.Duration(duration) * time.Second)
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

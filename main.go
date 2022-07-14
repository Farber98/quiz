package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

/*

REQUIREMENTS:
* Read quiz provided by CSV.
* Give the quiz to a user.
* Keep track of how many questions are right/incorrect
* Show next question inmediatly after answer.
* User can customize filename via flag
* Add customizable timer via flag. Quiz must stop when time limit has exceeded.
* User should be asked to press a key before timer starts.
* At the end of the timer, print total correct/total questions.

CONSTRAINTS:
* A quiz has < 100 questions
* Single word/number answers.
* Invalid answers are considered incorrect.

BONUS:
* Quiz shuffle.

*/

func init() {
	rand.Seed(time.Now().UnixNano())
}

func readCsvFile(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Couldn't read input file %s: %v", filepath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Couldn't parse file %s as CSV: %v", filepath, err)
	}

	return records
}

func shuffle(q [][]string) [][]string {
	rand.Shuffle(len(q), func(i, j int) { q[i], q[j] = q[j], q[i] })
	return q
}

func main() {
	file := flag.String("file", "problems.csv", ".csv filename to read from")
	timer := flag.Int("timer", 30, "The deadline for the countdown timer in seconds")
	flag.Parse()

	questions := shuffle(readCsvFile(*file))
	correct := 0

	var i string
	fmt.Println("Press ENTER to start the timer")
	fmt.Scanf("%s", &i)
	fmt.Printf("GO! You have %d seconds \n", *timer)

	deadline := time.NewTimer(time.Second * time.Duration(*timer))
	go func() {
		select {
		case <-deadline.C:
			fmt.Printf("\nOut of time! Your %d seconds timer finished.\n", *timer)
			fmt.Printf("You've scored %d out of %d\n", correct, len(questions))
			os.Exit(0)
		}
	}()

	for _, elem := range questions {
		var i string

		fmt.Printf("What is %s?: ", elem[0])

		fmt.Scanf("%s", &i)

		if i == elem[1] {
			correct++
		}
	}
	fmt.Printf("You've scored %d out of %d\n", correct, len(questions))
}

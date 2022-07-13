package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

/*

REQUIREMENTS:
* Read quiz provided by CSV.
* Give the quiz to a user.
* Keep track of how many questions are right/incorrect
* Show next question inmediatly after answer.
* User can customize filename via flag
* At the end, print total correct/total questions.

CONSTRAINTS:
* A quiz has < 100 questions
* Single word/number answers.
* Invalid answers are considered incorrect.


*/

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

func main() {
	var file string
	flag.StringVar(&file, "file", "problems.csv", ".csv filename to read from")
	flag.Parse()

	questions := readCsvFile(file)
	correct := 0

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

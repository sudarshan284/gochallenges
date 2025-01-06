/*

Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect
*/

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const totalQuestions = 5

type Question struct {
	question string
	answer   string
}

func main() {

}

func readArguments() (string, int) {
	filename := flag.String("filename", "problems.csv", "CSV files that contains quiz")
	timeLimit := flag.Int("Limit", 30, "Time limit for each question")
	flag.Parse()
	return *filename, *timeLimit
}

func readCSV(f io.Reader) ([]Question, error) {
	allQuestions, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	numofQues := len(allQuestions)
	if numofQues == 0 {
		return nil, fmt.Errorf("No questio in file")
	}

	var data []Question

	for _, line := range allQuestions {
		ques := Question{}
		ques.question = line[0]
		ques.answer = line[1]
		data = append(data, ques)
	}
	return data, nil
}

func openFile(filename string) (io.Reader, error) {
	return os.Open(filename)
}

func getInput(input chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input <- result
	}
}

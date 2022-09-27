package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type CsvDataLines struct {
	Question string
	Answer   string
}

func ReadCsvFile(filename string) ([][]string, error) {
	// Open CSV file
	fileContent, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer fileContent.Close()
	// Read File into a Variable
	lines, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func printResults(p int) {
	fmt.Println("Your points are ", p)
	fmt.Println("Thankyou!")
}

func main() {
	csvFileName := flag.String("csvFile", "problems.csv", "a csv file for problems which is in format question,answer")
	timeOut := flag.Duration("limit", 30, "a time limit for user to answer the questions ")
	flag.Parse()

	csvData, err := ReadCsvFile(*csvFileName)
	if err != nil {
		panic(err)
	}
	points := 0
	totalNumberOfQuestions := len(csvData)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to quiz game!!! Please enter you name: ")
	userName, _ := reader.ReadString('\n')
	fmt.Println("Hello ", userName, "There are ", totalNumberOfQuestions, "\nGame starts now !!")

	maxTimeOut := *timeOut
	expire := time.After(maxTimeOut * time.Second)
	go func() {
		if v := <-expire; v != nil {
			printResults(points)
			os.Exit(0)
		}
		select {
		case <-expire:

		}
	}()

	for _, line := range csvData {
		data := CsvDataLines{
			Question: line[0],
			Answer:   line[1],
		}
		fmt.Println(data.Question)
		userResponse, _ := reader.ReadString('\n')
		userResponse = strings.Replace(userResponse, "\n", "", -1)
		userResponse = strings.Trim(userResponse, " ")
		if userResponse == data.Answer {
			points++
		}
	}
	printResults(points)

}

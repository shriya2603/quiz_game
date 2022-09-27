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

type Problem struct {
	Question string
	Answer   string
}

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{
			Question: line[0],
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func printResults(points int, totalPoints int) {
	fmt.Printf("You scored %d out of %d.", points, totalPoints)
	fmt.Println("Thankyou!!")
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file for problems which is in format question,answer")
	timeOut := flag.Duration("limit", 30, "a time limit for user to answer the questions ")
	flag.Parse()

	fileContent, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file %s, error is %v", *csvFileName, err))
	}
	defer fileContent.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to Parse CSV file, error is %v", err))
	}
	points := 0
	totalNumberOfQuestions := len(lines)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to quiz game!!! Please enter you name: ")
	userName, _ := reader.ReadString('\n')
	fmt.Println("Hello ", userName, "There are ", totalNumberOfQuestions, "\nGame starts now !!")

	maxTimeOut := *timeOut
	expire := time.After(maxTimeOut * time.Second)
	go func() {
		select {
		case <-expire:
			printResults(points, len(lines))
			os.Exit(0)
		}
	}()

	problems := parseLines(lines)
	for i, problem := range problems {
		fmt.Printf("Problem #%d : %s =\n", i+1, problem.Question)
		var userResponse string
		fmt.Scanf("%s\n", &userResponse)
		if userResponse == problem.Answer {
			points++
		}
	}
	printResults(points, len(problems))

}
